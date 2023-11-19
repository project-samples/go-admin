package app

import (
	"context"
	"reflect"

	"github.com/core-go/auth"
	ah "github.com/core-go/auth/handler"
	. "github.com/core-go/auth/mock"
	as "github.com/core-go/auth/sql"
	"github.com/core-go/core/authorization"
	"github.com/core-go/core/code"
	. "github.com/core-go/core/health"
	. "github.com/core-go/core/jwt"
	. "github.com/core-go/core/security/sql"
	"github.com/core-go/core/shortid"
	"github.com/core-go/core/unique"
	v10 "github.com/core-go/core/v10"
	log "github.com/core-go/log/zap"
	"github.com/core-go/search/convert"
	"github.com/core-go/search/query"
	. "github.com/core-go/security"
	q "github.com/core-go/sql"
	sa "github.com/core-go/sql/action"
	"github.com/core-go/sql/template"
	"github.com/core-go/sql/template/xml"

	"go-service/internal/audit-log"
	r "go-service/internal/role"
	u "go-service/internal/user"
	"go-service/pkg/text"
)

type ApplicationContext struct {
	SkipSecurity         bool
	Health               *Handler
	Authorization        *authorization.Handler
	AuthorizationChecker *AuthorizationChecker
	Authorizer           *Authorizer
	Authentication       *ah.AuthenticationHandler
	Privileges           *ah.PrivilegesHandler
	Code                 *code.Handler
	Roles                *code.Handler
	Role                 r.RoleTransport
	User                 u.UserTransport
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
		logWriter := sa.NewActionLogWriter(auditLogDB, "auditlog", conf.AuditLog.Config, conf.AuditLog.Schema, generateId)
		writeLog = logWriter.Write
		auditLogHealthChecker := q.NewSqlHealthChecker(auditLogDB, "audit_log")
		healthHandler = NewHandler(sqlHealthChecker, auditLogHealthChecker)
	} else {
		healthHandler = NewHandler(sqlHealthChecker)
	}
	buildParam := q.GetBuild(db)
	validator, err := v10.NewValidator()
	if err != nil {
		return nil, err
	}
	sqlPrivilegeLoader := NewPrivilegeLoader(db, conf.Sql.PermissionsByUser)

	userId := conf.Tracking.User
	tokenService := NewTokenService()
	authorizationHandler := authorization.NewHandler(tokenService.GetAndVerifyToken, conf.Auth.Token.Secret)
	authorizationChecker := NewDefaultAuthorizationChecker(tokenService.GetAndVerifyToken, conf.Auth.Token.Secret, userId)
	authorizer := NewAuthorizer(sqlPrivilegeLoader.Privilege, true, userId)

	authStatus := auth.InitStatus(conf.Auth.Status)
	ldapAuthenticator, er2 := NewDAPAuthenticatorByConfig(conf.Ldap, authStatus)
	if er2 != nil {
		return nil, er2
	}
	userInfoService, er3 := as.NewUserRepository(db, conf.Auth.Query, conf.Auth.DB, conf.Auth.UserStatus)
	if er3 != nil {
		return nil, er3
	}
	privilegeLoader, er4 := as.NewSqlPrivilegesLoader(db, conf.Sql.PrivilegesByUser, 1, true)
	if er4 != nil {
		return nil, er4
	}
	authenticator := auth.NewBasicAuthenticator(authStatus, ldapAuthenticator.Authenticate, userInfoService, tokenService.GenerateToken, conf.Auth.Token, conf.Auth.Payload, privilegeLoader.Load)
	authenticationHandler := ah.NewAuthenticationHandler(authenticator.Authenticate, authStatus.Error, authStatus.Timeout, logError, writeLog)

	privilegeReader, er5 := as.NewPrivilegesReader(db, conf.Sql.Privileges)
	if er5 != nil {
		return nil, er5
	}
	privilegeHandler := ah.NewPrivilegesHandler(privilegeReader.Privileges)

	// codeLoader := code.NewDynamicSqlCodeLoader(db, "select code, name, status as text from codeMaster where master = ? and status = 'A'", 1)
	codeLoader, err := code.NewSqlCodeLoader(db, "codeMaster", conf.Code.Loader)
	if err != nil {
		return nil, err
	}
	codeHandler := code.NewCodeHandlerByConfig(codeLoader.Load, conf.Code.Handler, logError)

	templates, err := template.LoadTemplates(xml.Trim, "configs/query.xml")
	if err != nil {
		return nil, err
	}
	// rolesLoader, err := code.NewDynamicSqlCodeLoader(db, "select roleName as name, roleId as id from roles where status = 'A'", 0)

	rolesLoader, err := code.NewSqlCodeLoader(db, "roles", conf.Role.Loader)
	if err != nil {
		return nil, err
	}
	rolesHandler := code.NewCodeHandlerByConfig(rolesLoader.Load, conf.Role.Handler, logError)

	roleType := reflect.TypeOf(r.Role{})
	queryRole, err := template.GetQuery(conf.Template, query.UseQuery(db, "roles", roleType, buildParam), "role", templates, &roleType, convert.ToMap, buildParam, q.GetSort)
	roleSearchBuilder, err := q.NewSearchBuilder(db, roleType, queryRole)
	// roleValidator := user.NewRoleValidator(db, conf.Sql.Role.Duplicate, validator.validateFileName)
	roleValidator := unique.NewUniqueFieldValidator(db, "roles", "rolename", reflect.TypeOf(r.Role{}), validator.Validate)
	roleRepository, er6 := r.NewRoleAdapter(db, conf.Sql.Role.Check)
	if er6 != nil {
		return nil, er6
	}
	roleService := r.NewRoleService(roleRepository)
	generateRoleId := shortid.Func(conf.AutoRoleId)
	roleHandler := r.NewRoleHandler(roleSearchBuilder.Search, roleService, conf.Writer, logError, generateRoleId, roleValidator.Validate, conf.Tracking, writeLog)

	userType := reflect.TypeOf(u.User{})
	queryUser, err := template.GetQuery(conf.Template, query.UseQuery(db, "users", userType, buildParam), "user", templates, &userType, convert.ToMap, buildParam, q.GetSort)
	if err != nil {
		return nil, err
	}
	userSearchBuilder, err := q.NewSearchBuilder(db, userType, queryUser)
	if err != nil {
		return nil, err
	}
	// userValidator := user.NewUserValidator(db, conf.Sql.User, validator.validateFileName)
	userValidator := unique.NewUniqueFieldValidator(db, "users", "username", reflect.TypeOf(u.User{}), validator.Validate)
	userRepository, er7 := u.NewUserRepository(db)
	if er7 != nil {
		return nil, er7
	}
	userService := u.NewUserService(userRepository)
	generateUserId := shortid.Func(conf.AutoUserId)
	userHandler := u.NewUserHandler(userSearchBuilder.Search, userService, conf.Writer, logError, generateUserId, userValidator.Validate, conf.Tracking, writeLog)

	reportDB, er8 := q.Open(conf.AuditLog.DB)
	if er8 != nil {
		return nil, er8
	}
	userQuery, er := text.NewTextAdapter(db, "users", "userid", "email")
	if er != nil {
		return nil, er
	}

	auditLogQuery, er9 := audit.NewAuditLogQuery(reportDB, templates, userQuery.Query)
	if er9 != nil {
		return nil, er9
	}
	auditLogHandler := audit.NewAuditLogHandler(auditLogQuery, logError)

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
