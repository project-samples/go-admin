package app

import (
	"context"
	"net/http"

	c "github.com/core-go/core/constants"
	s "github.com/core-go/core/security"
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
	r.Use(app.Authorization.HandleAuthorization)
	sec := &s.SecurityConfig{SecuritySkip: conf.SecuritySkip, Check: app.AuthorizationChecker.Check, Authorize: app.Authorizer.Authorize}

	Handle(r, "/health", app.Health.Check, c.GET)
	Handle(r, "/authenticate", app.Authentication.Authenticate, c.POST)

	r.Handle("/code/{code}", app.AuthorizationChecker.Check(http.HandlerFunc(app.Code.Load))).Methods(c.GET)

	HandleWithSecurity(sec, r, "/privileges", app.Privileges.All, role, c.ActionRead, c.GET)
	roles := r.PathPrefix("/roles").Subrouter()
	HandleWithSecurity(sec, roles, "/search", app.Role.Search, role, c.ActionRead, c.POST, c.GET)
	HandleWithSecurity(sec, roles, "/{roleId}", app.Role.Load, role, c.ActionRead, c.GET)
	HandleWithSecurity(sec, roles, "", app.Role.Create, role, c.ActionWrite, c.POST)
	HandleWithSecurity(sec, roles, "/{roleId}", app.Role.Update, role, c.ActionWrite, c.PUT)
	HandleWithSecurity(sec, roles, "/{userId}", app.Role.Patch, user, c.ActionWrite, c.PATCH)
	HandleWithSecurity(sec, roles, "/{roleId}", app.Role.Delete, role, c.ActionWrite, c.DELETE)

	HandleWithSecurity(sec, r, "/roles", app.Roles.Load, user, c.ActionRead, c.GET)
	users := r.PathPrefix("/users").Subrouter()
	HandleWithSecurity(sec, users, "", app.User.GetUserByRole, role, c.ActionRead, c.GET)

	HandleWithSecurity(sec, users, "/search", app.User.Search, user, c.ActionRead, c.GET, c.POST)
	HandleWithSecurity(sec, users, "/{userId}", app.User.Load, user, c.ActionRead, c.GET)
	HandleWithSecurity(sec, users, "", app.User.Create, user, c.ActionWrite, c.POST)
	HandleWithSecurity(sec, users, "/{userId}", app.User.Update, user, c.ActionWrite, c.PUT)
	HandleWithSecurity(sec, users, "/{userId}", app.User.Patch, user, c.ActionWrite, c.PATCH)
	HandleWithSecurity(sec, users, "/{userId}", app.User.Delete, user, c.ActionWrite, c.DELETE)

	currency := "/currencies"
	r.HandleFunc(currency, app.Currency.Search).Methods(c.GET)
	r.HandleFunc(currency+"/search", app.Currency.Search).Methods(c.GET, c.POST)
	r.HandleFunc(currency+"/{id}", app.Currency.Load).Methods(c.GET)
	r.HandleFunc(currency, app.Currency.Create).Methods(c.POST)
	r.HandleFunc(currency+"/{id}", app.Currency.Update).Methods(c.PUT)
	r.HandleFunc(currency+"/{id}", app.Currency.Patch).Methods(c.PATCH)
	r.HandleFunc(currency+"/{id}", app.Currency.Delete).Methods(c.DELETE)

	locale := "/locales"
	r.HandleFunc(locale, app.Locale.Search).Methods(c.GET)
	r.HandleFunc(locale+"/search", app.Locale.Search).Methods(c.GET, c.POST)
	r.HandleFunc(locale+"/{id}", app.Locale.Load).Methods(c.GET)
	r.HandleFunc(locale, app.Locale.Create).Methods(c.POST)
	r.HandleFunc(locale+"/{id}", app.Locale.Update).Methods(c.PUT)
	r.HandleFunc(locale+"/{id}", app.Locale.Patch).Methods(c.PATCH)
	r.HandleFunc(locale+"/{id}", app.Locale.Delete).Methods(c.DELETE)

	HandleWithSecurity(sec, r, "/audit-logs", app.AuditLog.Search, audit_log, c.ActionRead, c.GET, c.POST)
	HandleWithSecurity(sec, r, "/audit-logs/search", app.AuditLog.Search, audit_log, c.ActionRead, c.GET, c.POST)
	Handle(r, "/settings", app.Settings.Save, c.PATCH)
	return nil
}

func Handle(r *mux.Router, path string, f func(http.ResponseWriter, *http.Request), methods ...string) *mux.Route {
	return r.HandleFunc(path, f).Methods(methods...)
}
func HandleWithSecurity(authorizer *s.SecurityConfig, r *mux.Router, path string, f func(http.ResponseWriter, *http.Request), menuId string, action int32, methods ...string) *mux.Route {
	finalHandler := http.HandlerFunc(f)
	if authorizer.SecuritySkip {
		return r.HandleFunc(path, finalHandler).Methods(methods...)
	}
	authorize := func(next http.Handler) http.Handler {
		return authorizer.Authorize(next, menuId, action)
	}
	return r.Handle(path, authorizer.Check(authorize(finalHandler))).Methods(methods...)
}
