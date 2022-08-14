package app

import (
	"context"
	"reflect"

	"github.com/core-go/auth"
	. "github.com/core-go/auth/mock"
	as "github.com/core-go/auth/sql"
	"github.com/core-go/code"
	"github.com/core-go/core/authorization"
	"github.com/core-go/core/shortid"
	"github.com/core-go/core/unique"
	v10 "github.com/core-go/core/v10"
	. "github.com/core-go/health"
	"github.com/core-go/log/zap"
	"github.com/core-go/search/convert"
	"github.com/core-go/search/query"
	"github.com/core-go/search/template"
	. "github.com/core-go/security"
	. "github.com/core-go/security/jwt"
	. "github.com/core-go/security/sql"
	q "github.com/core-go/sql"
	_ "github.com/go-sql-driver/mysql"

	"go-service/internal/usecase/audit-log"
	r "go-service/internal/usecase/role"
	u "go-service/internal/usecase/user"
)

type ApplicationContext struct {
	SkipSecurity         bool
	Health               *Handler
	Authorization        *authorization.Handler
	AuthorizationChecker *AuthorizationChecker
	Authorizer           *Authorizer
	Authentication       *auth.AuthenticationHandler
	Privileges           *auth.PrivilegesHandler
	Code                 *code.Handler
	Roles                *code.Handler
	Role                 r.RoleHandler
	User                 u.UserHandler
	AuditLog             *audit.AuditLogHandler
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	db, er0 := q.Open(conf.DB)
	if er0 != nil {
		return nil, er0
	}
	sqlHealthChecker := q.NewHealthChecker(db)
	var healthHandler *Handler

	logError := log.LogError
	generateId := shortid.Generate
	var writeLog func(ctx context.Context, resource string, action string, success bool, desc string) error

	if conf.AuditLog.Log {
		auditLogDB, er1 := q.Open(conf.AuditLog.DB)
		if er1 != nil {
			return nil, er1
		}
		logWriter := q.NewActionLogWriter(auditLogDB, "auditLog", conf.AuditLog.Config, conf.AuditLog.Schema, generateId)
		writeLog = logWriter.Write
		auditLogHealthChecker := q.NewSqlHealthChecker(auditLogDB, "audit_log")
		healthHandler = NewHandler(sqlHealthChecker, auditLogHealthChecker)
	} else {
		healthHandler = NewHandler(sqlHealthChecker)
	}
	buildParam := q.GetBuild(db)
	validator := v10.NewValidator()
	sqlPrivilegeLoader := NewPrivilegeLoader(db, conf.Sql.PermissionsByUser)

	userId := conf.Tracking.User
	tokenService := NewTokenService()
	authorizationHandler := authorization.NewHandler(tokenService.GetAndVerifyToken, conf.Token.Secret)
	authorizationChecker := NewDefaultAuthorizationChecker(tokenService.GetAndVerifyToken, conf.Token.Secret, userId)
	authorizer := NewAuthorizer(sqlPrivilegeLoader.Privilege, true, userId)

	authStatus := auth.InitStatus(conf.Status)
	ldapAuthenticator, er2 := NewDAPAuthenticatorByConfig(conf.Ldap, authStatus)
	if er2 != nil {
		return nil, er2
	}
	userInfoService, er3 := as.NewSqlUserInfoByConfig(db, conf.Auth)
	if er3 != nil {
		return nil, er3
	}
	privilegeLoader, er4 := as.NewSqlPrivilegesLoader(db, conf.Sql.PrivilegesByUser, 1, true)
	if er4 != nil {
		return nil, er4
	}
	authenticator := auth.NewBasicAuthenticator(authStatus, ldapAuthenticator.Authenticate, userInfoService, tokenService.GenerateToken, conf.Token, conf.Payload, privilegeLoader.Load)
	authenticationHandler := auth.NewAuthenticationHandler(authenticator.Authenticate, authStatus.Error, authStatus.Timeout, log.ErrorMsg, writeLog)

	privilegeReader, er5 := as.NewPrivilegesReader(db, conf.Sql.Privileges)
	if er5 != nil {
		return nil, er5
	}
	privilegeHandler := auth.NewPrivilegesHandler(privilegeReader.Privileges)

	// codeLoader := code.NewDynamicSqlCodeLoader(db, "select code, name, status as text from codeMaster where master = ? and status = 'A'", 1)
	codeLoader := code.NewSqlCodeLoader(db, "codeMaster", conf.Code.Loader)
	codeHandler := code.NewCodeHandlerByConfig(codeLoader.Load, conf.Code.Handler, log.ErrorMsg)

	templates, err := template.LoadTemplates(template.Trim, "configs/query.xml")
	if err != nil {
		return nil, err
	}
	// rolesLoader := code.NewDynamicSqlCodeLoader(db, "select roleName as name, roleId as id, status as code from roles where status = 'A'", 0)
	rolesLoader := code.NewSqlCodeLoader(db, "roles", conf.Role.Loader)
	rolesHandler := code.NewCodeHandlerByConfig(rolesLoader.Load, conf.Role.Handler, log.ErrorMsg)

	roleType := reflect.TypeOf(r.Role{})
	queryRole, err := template.UseQuery(conf.Template, query.UseQuery(db, "roles", roleType, buildParam), "role", templates, &roleType, convert.ToMap, buildParam)
	roleSearchBuilder, err := q.NewSearchBuilder(db, roleType, queryRole)
	// roleValidator := user.NewRoleValidator(db, conf.Sql.Role.Duplicate, validator.Validate)
	roleValidator := unique.NewUniqueFieldValidator(db, "roles", "rolename", reflect.TypeOf(r.Role{}), validator.Validate)
	roleService, er6 := r.NewRoleService(db, conf.Sql.Role.Check)
	if er6 != nil {
		return nil, er6
	}
	generateRoleId := shortid.Func(conf.AutoRoleId)
	roleHandler := r.NewRoleHandler(roleSearchBuilder.Search, roleService, conf.Writer, logError, generateRoleId, roleValidator.Validate, conf.Tracking, writeLog)

	userType := reflect.TypeOf(u.User{})
	queryUser, err := template.UseQuery(conf.Template, query.UseQuery(db, "users", userType, buildParam), "user", templates, &userType, convert.ToMap, buildParam)
	if err != nil {
		return nil, err
	}
	userSearchBuilder, err := q.NewSearchBuilder(db, userType, queryUser)
	if err != nil {
		return nil, err
	}
	// userValidator := user.NewUserValidator(db, conf.Sql.User, validator.Validate)
	userValidator := unique.NewUniqueFieldValidator(db, "users", "username", reflect.TypeOf(u.User{}), validator.Validate)
	userService, er7 := u.NewUserService(db)
	if er7 != nil {
		return nil, er7
	}
	generateUserId := shortid.Func(conf.AutoUserId)
	userHandler := u.NewUserHandler(userSearchBuilder.Search, userService, conf.Writer, logError, generateUserId, userValidator.Validate, conf.Tracking, writeLog)

	reportDB, er8 := q.Open(conf.AuditLog.DB)
	if er8 != nil {
		return nil, er8
	}
	auditLogService, er9 := audit.NewAuditLogService(reportDB)
	if er9 != nil {
		return nil, er9
	}
	auditLogHandler := audit.NewAuditLogHandler(auditLogService, logError, writeLog)

	app := &ApplicationContext{
		Health:               healthHandler,
		SkipSecurity:         conf.SecuritySkip,
		Authorization:        authorizationHandler,
		AuthorizationChecker: authorizationChecker,
		Authorizer:           authorizer,
		Authentication:       authenticationHandler,
		Privileges:           privilegeHandler,
		Code:                 codeHandler,
		Roles:                rolesHandler,
		Role:                 roleHandler,
		User:                 userHandler,
		AuditLog:             auditLogHandler,
	}
	return app, nil
}
