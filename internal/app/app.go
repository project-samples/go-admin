package app

import (
	"context"
	"database/sql"
	auth "github.com/core-go/authentication"
	ah "github.com/core-go/authentication/handler"
	"github.com/core-go/authentication/mock"
	as "github.com/core-go/authentication/sql"
	"github.com/core-go/core/authorization"
	"github.com/core-go/core/code"
	hs "github.com/core-go/core/health/sql"
	"github.com/core-go/core/jwt"
	se "github.com/core-go/core/settings"
	"github.com/core-go/core/shortid"
	ur "github.com/core-go/core/user"
	"github.com/core-go/health"
	log "github.com/core-go/log/zap"
	"github.com/core-go/redis/v9"
	sec "github.com/core-go/security"
	"github.com/core-go/security/cookies"
	ss "github.com/core-go/security/sql"
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
	db, er0 := sql.Open(cfg.DB.Driver, cfg.DB.DataSourceName)
	if er0 != nil {
		return nil, er0
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	sqlHealthChecker := hs.NewHealthChecker(db)
	var healthHandler *health.Handler
	cachePort, err := redis.NewRedisAdapterByConfig(cfg.Redis)
	/*
		if err != nil {
			return nil, err
		}
		{
			err := cachePort.Pool.Ping(ctx).Err()
			if err != nil {
				return nil, err
			}
		}*/
	redisChecker := redis.NewHealthChecker(cachePort.Client)
	logError := log.LogError
	generateId := shortid.Generate
	var writeLog func(ctx context.Context, resource string, action string, success bool, desc string) error

	if cfg.AuditLog.Log {
		auditLogDB, er1 := sql.Open(cfg.AuditLog.DB.Driver, cfg.AuditLog.DB.DataSourceName)
		if er1 != nil {
			return nil, er1
		}
		logWriter := sa.NewActionLogWriter(auditLogDB, "audit_logs", cfg.AuditLog.Config, cfg.AuditLog.Schema, generateId)
		writeLog = logWriter.Write
		auditLogHealthChecker := hs.NewSqlHealthChecker(auditLogDB, "audit_log")
		healthHandler = health.NewHandler(sqlHealthChecker, auditLogHealthChecker, redisChecker)
	} else {
		healthHandler = health.NewHandler(sqlHealthChecker, redisChecker)
	}

	sqlPrivilegeLoader := ss.NewPrivilegeLoader(db, cfg.Sql.PermissionsByUser)

	userId := cfg.Tracking.User
	tokenPort := jwt.NewTokenAdapter()
	authorizationHandler := authorization.NewHandler(tokenPort.GetAndVerifyToken, cfg.Auth.Token.Secret)
	authorizationChecker := sec.NewAuthorizationChecker(tokenPort.GetAndVerifyToken, cfg.Auth.Token.Secret, userId)
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
	codeLoader, err := code.NewSqlCodeLoader(db, "code_masters", cfg.Code.Loader)
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

	roleHandler, err := r.NewRoleTransport(db, logError, templates, cfg.Tracking, writeLog, cfg.Action)
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

	localeHandler, err := loc.NewLocaleTransport(db, logError, writeLog, nil)
	if err != nil {
		return nil, err
	}

	countryHandler, err := country.NewCountryTransport(db, logError, writeLog, nil)
	if err != nil {
		return nil, err
	}

	reportDB, er8 := sql.Open(cfg.AuditLog.DB.Driver, cfg.AuditLog.DB.DataSourceName)
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
