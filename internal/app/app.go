package app

import (
	"context"
	role2 "go-service/internal/role"
	"reflect"

	"github.com/core-go/auth"
	. "github.com/core-go/auth/mock"
	as "github.com/core-go/auth/sql"
	"github.com/core-go/code"
	. "github.com/core-go/health"
	"github.com/core-go/log/zap"
	. "github.com/core-go/security"
	. "github.com/core-go/security/jwt"
	. "github.com/core-go/security/sql"
	sv "github.com/core-go/service"
	"github.com/core-go/service/shortid"
	"github.com/core-go/service/unique"
	v10 "github.com/core-go/service/v10"
	s "github.com/core-go/sql"
	_ "github.com/go-sql-driver/mysql"

	"go-service/internal/uam"
)

type ApplicationContext struct {
	SkipSecurity          bool
	HealthHandler         *Handler
	AuthorizationHandler  *sv.AuthorizationHandler
	AuthorizationChecker  *AuthorizationChecker
	Authorizer            *Authorizer
	AuthenticationHandler *auth.AuthenticationHandler
	PrivilegesHandler     *auth.PrivilegesHandler
	CodeHandler           *code.Handler
	RolesHandler          *code.Handler
	RoleHandler           *role2.RoleHandler
	UserHandler           *uam.UserHandler
}

func NewApp(ctx context.Context, conf Root) (*ApplicationContext, error) {
	db, er1 := s.Open(conf.DB)
	if er1 != nil {
		return nil, er1
	}
	sqlHealthChecker := s.NewHealthChecker(db)
	var healthHandler *Handler

	logError := log.ErrorMsg
	generateId := shortid.Generate
	var writeLog func(ctx context.Context, resource string, action string, success bool, desc string) error

	if conf.AuditLog.Log {
		auditLogDB, er2 := s.Open(conf.AuditLog.DB)
		if er2 != nil {
			return nil, er2
		}
		logWriter := s.NewActionLogWriter(auditLogDB, "auditLog", conf.AuditLog.Config, conf.AuditLog.Schema, generateId)
		writeLog = logWriter.Write
		auditLogHealthChecker := s.NewSqlHealthChecker(auditLogDB, "audit_log")
		healthHandler = NewHandler(sqlHealthChecker, auditLogHealthChecker)
	} else {
		healthHandler = NewHandler(sqlHealthChecker)
	}

	validator := v10.NewValidator()
	sqlPrivilegeLoader := NewPrivilegeLoader(db, conf.Sql.PermissionsByUser)

	userId := conf.Tracking.User
	tokenService := NewTokenService()
	authorizationHandler := sv.NewAuthorizationHandler(tokenService.GetAndVerifyToken, conf.Token.Secret)
	authorizationChecker := NewDefaultAuthorizationChecker(tokenService.GetAndVerifyToken, conf.Token.Secret, userId)
	authorizer := NewAuthorizer(sqlPrivilegeLoader.Privilege, true, userId)

	authStatus := auth.InitStatus(conf.Status)
	ldapAuthenticator, er3 := NewDAPAuthenticatorByConfig(conf.Ldap, authStatus)
	if er3 != nil {
		return nil, er3
	}
	userInfoService := as.NewSqlUserInfoByConfig(db, conf.Auth)
	privilegeLoader := as.NewSqlPrivilegesLoader(db, conf.Sql.PrivilegesByUser, 1, true)
	authenticator := auth.NewBasicAuthenticator(authStatus, ldapAuthenticator.Authenticate, userInfoService, tokenService.GenerateToken, conf.Token, conf.Payload, privilegeLoader.Load)
	authenticationHandler := auth.NewAuthenticationHandler(authenticator.Authenticate, authStatus.Error, authStatus.Timeout, logError, writeLog)

	privilegeReader := as.NewPrivilegesReader(db, conf.Sql.Privileges)
	privilegeHandler := auth.NewPrivilegesHandler(privilegeReader.Privileges)

	// codeLoader := code.NewDynamicSqlCodeLoader(db, "select code, name, status as text from codeMaster where master = ? and status = 'A'", 1)
	codeLoader := code.NewSqlCodeLoader(db, "codeMaster", conf.Code.Loader)
	codeHandler := code.NewCodeHandlerByConfig(codeLoader.Load, conf.Code.Handler, logError)

	// rolesLoader := code.NewDynamicSqlCodeLoader(db, "select roleName as name, roleId as id, status as code from roles where status = 'A'", 0)
	rolesLoader := code.NewSqlCodeLoader(db, "roles", conf.Role.Loader)
	rolesHandler := code.NewCodeHandlerByConfig(rolesLoader.Load, conf.Role.Handler, logError)

	roleService := role2.NewRoleService(db, conf.Sql.Role.Check)
	roleValidator := unique.NewUniqueFieldValidator(db, "roles", "rolename", reflect.TypeOf(role2.Role{}), validator.Validate)
	// roleValidator := uam.NewRoleValidator(db, conf.Sql.Role.Duplicate, validator.Validate)
	generateRoleId := shortid.Func(conf.AutoRoleId)
	roleHandler := role2.NewRoleHandler(roleService, conf.Writer, logError, generateRoleId, roleValidator.Validate, conf.Tracking, writeLog)

	userService := uam.NewUserService(db)
	userValidator := unique.NewUniqueFieldValidator(db, "users", "username", reflect.TypeOf(uam.User{}), validator.Validate)
	// userValidator := uam.NewUserValidator(db, conf.Sql.User, validator.Validate)
	generateUserId := shortid.Func(conf.AutoUserId)
	userHandler := uam.NewUserHandler(userService, conf.Writer, logError, generateUserId, userValidator.Validate, conf.Tracking, writeLog)

	app := &ApplicationContext{
		HealthHandler:         healthHandler,
		SkipSecurity:          conf.SecuritySkip,
		AuthorizationHandler:  authorizationHandler,
		AuthorizationChecker:  authorizationChecker,
		Authorizer:            authorizer,
		AuthenticationHandler: authenticationHandler,
		PrivilegesHandler:     privilegeHandler,
		CodeHandler:           codeHandler,
		RolesHandler:          rolesHandler,
		RoleHandler:           roleHandler,
		UserHandler:           userHandler,
	}
	return app, nil
}
