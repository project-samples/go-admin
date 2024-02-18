package app

import (
	"context"
	"reflect"

	"github.com/core-go/auth"
	ah "github.com/core-go/auth/handler"
	"github.com/core-go/auth/mock"
	as "github.com/core-go/auth/sql"
	"github.com/core-go/core"
	"github.com/core-go/core/authorization"
	"github.com/core-go/core/code"
	"github.com/core-go/core/health"
	"github.com/core-go/core/jwt"
	sec "github.com/core-go/core/security"
	ss "github.com/core-go/core/security/sql"
	se "github.com/core-go/core/settings"
	"github.com/core-go/core/shortid"
	"github.com/core-go/core/unique"
	ur "github.com/core-go/core/user"
	v10 "github.com/core-go/core/v10"
	log "github.com/core-go/log/zap"
	"github.com/core-go/search/convert"
	"github.com/core-go/search/query"
	q "github.com/core-go/sql"
	sa "github.com/core-go/sql/action"
	"github.com/core-go/sql/template"
	"github.com/core-go/sql/template/xml"

	"go-service/internal/audit-log"
	c "go-service/internal/currency"
	loc "go-service/internal/locale"
	r "go-service/internal/role"
	u "go-service/internal/user"
)

type ApplicationContext struct {
	SkipSecurity         bool
	Health               *health.Handler
	Authorization        *authorization.Handler
	AuthorizationChecker *sec.AuthorizationChecker
	Authorizer           *sec.Authorizer
	Authentication       *ah.AuthenticationHandler
	Privileges           *ah.PrivilegesHandler
	Code                 *code.Handler
	Roles                *code.Handler
	Role                 r.RoleTransport
	User                 u.UserTransport
	Currency             c.CurrencyTransport
	Locale               loc.LocaleTransport
	AuditLog             *audit.AuditLogHandler
	Settings             *se.Handler
}

func NewApp(ctx context.Context, cfg Config) (*ApplicationContext, error) {
	db, er0 := q.Open(cfg.DB)
	if er0 != nil {
		return nil, er0
	}
	sqlHealthChecker := q.NewHealthChecker(db)
	var healthHandler *health.Handler

	logError := log.LogError
	generateId := shortid.Generate
	var writeLog func(ctx context.Context, resource string, action string, success bool, desc string) error

	if cfg.AuditLog.Log {
		auditLogDB, er1 := q.Open(cfg.AuditLog.DB)
		if er1 != nil {
			return nil, er1
		}
		logWriter := sa.NewActionLogWriter(auditLogDB, "auditlog", cfg.AuditLog.Config, cfg.AuditLog.Schema, generateId)
		writeLog = logWriter.Write
		auditLogHealthChecker := q.NewSqlHealthChecker(auditLogDB, "audit_log")
		healthHandler = health.NewHandler(sqlHealthChecker, auditLogHealthChecker)
	} else {
		healthHandler = health.NewHandler(sqlHealthChecker)
	}
	buildParam := q.GetBuild(db)
	validator, err := v10.NewValidator()
	if err != nil {
		return nil, err
	}
	sqlPrivilegeLoader := ss.NewPrivilegeLoader(db, cfg.Sql.PermissionsByUser)

	userId := cfg.Tracking.User
	tokenService := jwt.NewTokenService()
	authorizationHandler := authorization.NewHandler(tokenService.GetAndVerifyToken, cfg.Auth.Token.Secret)
	authorizationChecker := sec.NewDefaultAuthorizationChecker(tokenService.GetAndVerifyToken, cfg.Auth.Token.Secret, userId)
	authorizer := sec.NewAuthorizer(sqlPrivilegeLoader.Privilege, true, userId)

	authStatus := auth.InitStatus(cfg.Auth.Status)
	ldapAuthenticator, er2 := mock.NewDAPAuthenticatorByConfig(cfg.Ldap, authStatus)
	if er2 != nil {
		return nil, er2
	}
	userInfoService, er3 := as.NewUserRepository(db, cfg.Auth.Query, cfg.Auth.DB, cfg.Auth.UserStatus)
	if er3 != nil {
		return nil, er3
	}
	privilegeLoader, er4 := as.NewSqlPrivilegesLoader(db, cfg.Sql.PrivilegesByUser, 1, true)
	if er4 != nil {
		return nil, er4
	}
	authenticator := auth.NewBasicAuthenticator(authStatus, ldapAuthenticator.Authenticate, userInfoService, tokenService.GenerateToken, cfg.Auth.Token, cfg.Auth.Payload, privilegeLoader.Load)
	authenticationHandler := ah.NewAuthenticationHandler(authenticator.Authenticate, authStatus.Error, authStatus.Timeout, logError, writeLog)
	authenticationHandler.Cookie = false

	privilegeReader, er5 := as.NewPrivilegesReader(db, cfg.Sql.Privileges)
	if er5 != nil {
		return nil, er5
	}
	privilegeHandler := ah.NewPrivilegesHandler(privilegeReader.Privileges)

	// codeLoader := code.NewDynamicSqlCodeLoader(db, "select code, name, status as text from codeMaster where master = ? and status = 'A'", 1)
	codeLoader, err := code.NewSqlCodeLoader(db, "code_master", cfg.Code.Loader)
	if err != nil {
		return nil, err
	}
	codeHandler := code.NewCodeHandlerByConfig(codeLoader.Load, cfg.Code.Handler, logError)

	templates, err := template.LoadTemplates(xml.Trim, "configs/query.xml")
	if err != nil {
		return nil, err
	}
	// rolesLoader, err := code.NewDynamicSqlCodeLoader(db, "select roleName as name, roleId as id from roles where status = 'A'", 0)

	rolesLoader, err := code.NewSqlCodeLoader(db, "roles", cfg.Role.Loader)
	if err != nil {
		return nil, err
	}
	rolesHandler := code.NewCodeHandlerByConfig(rolesLoader.Load, cfg.Role.Handler, logError)

	roleType := reflect.TypeOf(r.Role{})
	queryRole, err := template.GetQuery(cfg.Template, query.UseQuery(db, "roles", roleType, buildParam), "role", templates, &roleType, convert.ToMap, buildParam, q.GetSort)
	if err != nil {
		return nil, err
	}
	roleSearchBuilder, err := q.NewSearchBuilder(db, roleType, queryRole)
	if err != nil {
		return nil, err
	}
	// roleValidator := user.NewRoleValidator(db, conf.Sql.Role.Duplicate, validator.validateFileName)
	roleValidator := unique.NewUniqueFieldValidator(db, "roles", "rolename", reflect.TypeOf(r.Role{}), validator.Validate)
	roleRepository, er6 := r.NewRoleAdapter(db, cfg.Sql.Role.Check)
	if er6 != nil {
		return nil, er6
	}
	roleService := r.NewRoleService(roleRepository)
	generateRoleId := shortid.Func(cfg.AutoRoleId)
	roleHandler := r.NewRoleHandler(roleSearchBuilder.Search, roleService, generateRoleId, roleValidator.Validate, logError, writeLog, cfg.Action, cfg.Tracking)

	userType := reflect.TypeOf(u.User{})
	queryUser, err := template.GetQuery(cfg.Template, query.UseQuery(db, "users", userType, buildParam), "user", templates, &userType, convert.ToMap, buildParam, q.GetSort)
	if err != nil {
		return nil, err
	}
	userSearchBuilder, err := q.NewSearchBuilder(db, userType, queryUser)
	if err != nil {
		return nil, err
	}
	// userValidator := user.NewUserValidator(db, conf.Sql.User, validator.validateFileName)
	userValidator := unique.NewUniqueFieldValidator(db, "users", "username", reflect.TypeOf(u.User{}), validator.Validate)
	userRepository, er7 := u.NewUserAdapter(db)
	if er7 != nil {
		return nil, er7
	}
	userService := u.NewUserService(userRepository)
	generateUserId := shortid.Func(cfg.AutoUserId)
	userHandler := u.NewUserHandler(userSearchBuilder.Search, userService, generateUserId, userValidator.Validate, logError, writeLog, cfg.Action, cfg.Tracking)

	action := core.InitializeAction(cfg.Action)
	currencyType := reflect.TypeOf(c.Currency{})
	currencyQueryBuilder := query.UseQuery(db, "currency", currencyType)
	currencySearchBuilder, err := q.NewSearchBuilder(db, currencyType, currencyQueryBuilder)
	if err != nil {
		return nil, err
	}
	currencyRepository, err := q.NewRepository(db, "currency", currencyType)
	if err != nil {
		return nil, err
	}
	currencyService := c.NewCurrencyService(currencyRepository)
	currencyHandler := c.NewCurrencyHandler(currencySearchBuilder.Search, currencyService, logError, validator.Validate, &action)

	localeType := reflect.TypeOf(loc.Locale{})
	queryLocale := query.UseQuery(db, "locale", localeType)
	localeSearchBuilder, err := q.NewSearchBuilder(db, localeType, queryLocale)
	if err != nil {
		return nil, err
	}
	localeRepository, err := q.NewRepository(db, "locale", localeType)
	if err != nil {
		return nil, err
	}
	localeService := loc.NewLocaleService(localeRepository)
	localeHandler := loc.NewLocaleHandler(localeSearchBuilder.Search, localeService, logError, validator.Validate, &action)

	reportDB, er8 := q.Open(cfg.AuditLog.DB)
	if er8 != nil {
		return nil, er8
	}
	userQuery := ur.NewUserAdapter(db, "select userId, displayName, email, phone, imageURL from users where userId ")

	auditLogQuery, er9 := audit.NewAuditLogQuery(reportDB, templates, userQuery.Query)
	if er9 != nil {
		return nil, er9
	}
	auditLogHandler := audit.NewAuditLogHandler(auditLogQuery, logError)

	settingsHandler := se.NewSettingsHandler(logError, writeLog, db, "users", buildParam, "userId", "userid", "dateformat", "language")

	app := &ApplicationContext{
		Health:               healthHandler,
		SkipSecurity:         cfg.SecuritySkip,
		Authorization:        authorizationHandler,
		AuthorizationChecker: authorizationChecker,
		Authorizer:           authorizer,
		Authentication:       authenticationHandler,
		Privileges:           privilegeHandler,
		Code:                 codeHandler,
		Roles:                rolesHandler,
		Role:                 roleHandler,
		User:                 userHandler,
		Currency:             currencyHandler,
		Locale:               localeHandler,
		AuditLog:             auditLogHandler,
		Settings:             settingsHandler,
	}
	return app, nil
}
