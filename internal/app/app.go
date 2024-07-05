package app

import (
	"context"
	"github.com/core-go/auth"
	ah "github.com/core-go/auth/handler"
	"github.com/core-go/auth/mock"
	as "github.com/core-go/auth/sql"
	"github.com/core-go/core/authorization"
	"github.com/core-go/core/code"
	"github.com/core-go/core/cookies"
	"github.com/core-go/core/health"
	"github.com/core-go/core/jwt"
	redis "github.com/core-go/core/redis/v8"
	sec "github.com/core-go/core/security"
	ss "github.com/core-go/core/security/sql"
	se "github.com/core-go/core/settings"
	"github.com/core-go/core/shortid"
	ur "github.com/core-go/core/user"
	log "github.com/core-go/log/zap"
	q "github.com/core-go/sql"
	sa "github.com/core-go/sql/action"
	"github.com/core-go/sql/template"
	"github.com/core-go/sql/template/xml"
	"github.com/google/uuid"
	"net/http"

	"go-service/internal/audit-log"
	"go-service/internal/country"
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
	SessionAuthorizer    *sec.SessionAuthorizer
	Privileges           *ah.PrivilegesHandler
	Code                 *code.Handler
	Roles                *code.Handler
	Role                 r.RoleTransport
	User                 u.UserTransport
	Currency             c.CurrencyTransport
	Locale               loc.LocaleTransport
	Country              country.CountryTransport
	AuditLog             *audit.AuditLogHandler
	Settings             *se.Handler
}

func NewApp(ctx context.Context, cfg Config) (*ApplicationContext, error) {
	db, er0 := q.Open(cfg.DB)
	if er0 != nil {
		return nil, er0
	}
	if err := db.Ping(); err != nil {
		return nil, er0
	}
	sqlHealthChecker := q.NewHealthChecker(db)
	var healthHandler *health.Handler
	cachePort, err := redis.NewRedisAdapterByConfig(cfg.Redis)
	if err != nil {
		return nil, err
	}
	{
		err := cachePort.Client.Ping(ctx).Err()
		if err != nil {
			return nil, err
		}
	}
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

	sqlPrivilegeLoader := ss.NewPrivilegeLoader(db, cfg.Sql.PermissionsByUser)

	userId := cfg.Tracking.User
	tokenPort := jwt.NewTokenAdapter()
	authorizationHandler := authorization.NewHandler(tokenPort.GetAndVerifyToken, cfg.Auth.Token.Secret)
	authorizationChecker := sec.NewDefaultAuthorizationChecker(tokenPort.GetAndVerifyToken, cfg.Auth.Token.Secret, userId)
	authorizer := sec.NewAuthorizer(sqlPrivilegeLoader.Privilege, true, userId)

	authStatus := auth.InitStatus(cfg.Auth.Status)
	ldapAuthenticator, er2 := mock.NewDAPAuthenticatorByConfig(cfg.Ldap, authStatus)
	if er2 != nil {
		return nil, er2
	}
	userPort, er3 := as.NewUserAdapter(db, cfg.Auth.Query, cfg.Auth.DB, cfg.Auth.UserStatus)
	if er3 != nil {
		return nil, er3
	}
	privilegePort, er4 := as.NewSqlPrivilegesAdapter(db, cfg.Sql.PrivilegesByUser, 1, true)
	if er4 != nil {
		return nil, er4
	}
	generateUUID := func(ctx context.Context) (string, error) {
		return uuid.NewString(), nil
	}
	authenticator := auth.NewBasicAuthenticator(authStatus, ldapAuthenticator.Authenticate, userPort, tokenPort.GenerateToken, cfg.Auth.Token, cfg.Auth.Payload, privilegePort.Load)
	authenticationHandler := ah.NewAuthenticationHandlerWithCache(
		authenticator.Authenticate,
		authStatus.Error,
		authStatus.Timeout, logError, cachePort, generateUUID, cfg.Session.ExpiredTime, cfg.Session.Host, http.SameSiteStrictMode, true, true, writeLog)
	// jTokenSvc := jwt.NewTokenService()
	co := cookies.NewCookies("id", cfg.Session.Host, cfg.Session.ExpiredTime, http.SameSiteStrictMode)

	sessionAuthorizer := sec.NewSessionAuthorizer(cfg.Auth.Token.Secret, tokenPort.VerifyToken, co.RefreshValue, cachePort, cfg.Session.ExpiredTime, logError, true, nil, nil)

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

	roleHandler, err := r.NewRoleTransport(db, cfg.Sql.Role.Check, logError, templates, cfg.Tracking, writeLog, cfg.Action)
	if err != nil {
		return nil, err
	}

	userHandler, err := u.NewUserTransport(db, logError, templates, cfg.Tracking, writeLog, cfg.Action)
	if err != nil {
		return nil, err
	}

	currencyHandler, err := c.NewCurrencyTransport(db, logError, writeLog, nil)
	if err != nil {
		return nil, err
	}

	localeHandler, err := loc.NewLocaleTransport(db, logError, nil)
	if err != nil {
		return nil, err
	}

	countryHandler, err := country.NewCountryTransport(db, logError, writeLog, nil)
	if err != nil {
		return nil, err
	}

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

	buildParam := q.GetBuild(db)
	settingsHandler := se.NewSettingsHandler(logError, writeLog, db, "users", buildParam, "userId", "userid", "dateformat", "language")

	app := &ApplicationContext{
		Health:               healthHandler,
		SkipSecurity:         cfg.SecuritySkip,
		Authorization:        authorizationHandler,
		AuthorizationChecker: authorizationChecker,
		Authorizer:           authorizer,
		Authentication:       authenticationHandler,
		SessionAuthorizer:    sessionAuthorizer,
		Privileges:           privilegeHandler,
		Code:                 codeHandler,
		Roles:                rolesHandler,
		Role:                 roleHandler,
		User:                 userHandler,
		Currency:             currencyHandler,
		Locale:               localeHandler,
		Country:              countryHandler,
		AuditLog:             auditLogHandler,
		Settings:             settingsHandler,
	}
	return app, nil
}
