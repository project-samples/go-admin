package app

import (
	"context"
	"net/http"

	. "github.com/core-go/security"
	"github.com/gorilla/mux"
)

const (
	role = "role"
	user = "user"
	audit_log   = "audit-log"
)

func Route(r *mux.Router, ctx context.Context, conf Config) error {
	app, err := NewApp(ctx, conf)
	if err != nil {
		return err
	}

	r.Use(app.AuthorizationHandler.HandleAuthorization)
	sec := &SecurityConfig{SecuritySkip: conf.SecuritySkip, Check: app.AuthorizationChecker.Check, Authorize: app.Authorizer.Authorize}

	Handle(r, "/health", app.HealthHandler.Check, GET)
	Handle(r, "/authenticate", app.AuthenticationHandler.Authenticate, POST)

	r.Handle("/code/{code}", app.AuthorizationChecker.Check(http.HandlerFunc(app.CodeHandler.Load))).Methods(GET)

	HandleWithSecurity(sec, r, "/privileges", app.PrivilegesHandler.All, role, ActionRead, GET)
	roleHandler := app.RoleHandler
	roles := r.PathPrefix("/roles").Subrouter()
	HandleWithSecurity(sec, roles, "/search", roleHandler.Search, role, ActionRead, POST, GET)
	HandleWithSecurity(sec, roles, "/{roleId}", roleHandler.Load, role, ActionRead, GET)
	HandleWithSecurity(sec, roles, "", roleHandler.Create, role, ActionWrite, POST)
	HandleWithSecurity(sec, roles, "/{roleId}", roleHandler.Update, role, ActionWrite, PUT)
	HandleWithSecurity(sec, roles, "/{userId}", roleHandler.Patch, user, ActionWrite, PATCH)
	HandleWithSecurity(sec, roles, "/{roleId}", roleHandler.Delete, role, ActionWrite, DELETE)

	HandleWithSecurity(sec, r, "/roles", app.RolesHandler.Load, user, ActionRead, GET)
	userHandler := app.UserHandler
	users := r.PathPrefix("/users").Subrouter()
	HandleWithSecurity(sec, users, "", userHandler.Search, user, ActionRead, GET)
	HandleWithSecurity(sec, users, "/search", userHandler.Search, user, ActionRead, GET, POST)
	HandleWithSecurity(sec, users, "/{userId}", userHandler.Load, user, ActionRead, GET)
	HandleWithSecurity(sec, users, "", userHandler.Create, user, ActionWrite, POST)
	HandleWithSecurity(sec, users, "/{userId}", userHandler.Update, user, ActionWrite, PUT)
	HandleWithSecurity(sec, users, "/{userId}", userHandler.Patch, user, ActionWrite, PATCH)
	HandleWithSecurity(sec, users, "/{userId}", userHandler.Delete, user, ActionWrite, DELETE)

	HandleWithSecurity(sec, r, "/audit-logs", app.AuditLogHandler.Search, audit_log, ActionRead, GET, POST)
	HandleWithSecurity(sec, r, "/audit-logs/search", app.AuditLogHandler.Search, audit_log, ActionRead, GET, POST)
	return nil
}

func Handle(r *mux.Router, path string, f func(http.ResponseWriter, *http.Request), methods ...string) *mux.Route {
	return r.HandleFunc(path, f).Methods(methods...)
}
func HandleWithSecurity(authorizer *SecurityConfig, r *mux.Router, path string, f func(http.ResponseWriter, *http.Request), menuId string, action int32, methods ...string) *mux.Route {
	finalHandler := http.HandlerFunc(f)
	if authorizer.SecuritySkip {
		return r.HandleFunc(path, finalHandler).Methods(methods...)
	}
	authorize := func(next http.Handler) http.Handler {
		return authorizer.Authorize(next, menuId, action)
	}
	return r.Handle(path, authorizer.Check(authorize(finalHandler))).Methods(methods...)
}
