package app

import (
	"context"
	"net/http"

	c "github.com/core-go/core/constants"
	m "github.com/core-go/core/mux"
	s "github.com/core-go/core/security"
	"github.com/gorilla/mux"
)

const (
	role = "role"
	user = "user"
	currency = "currency"
	locale = "locale"
	Country = "country"
	audit_log = "audit_log"
)

func Route(r *mux.Router, ctx context.Context, conf Config) error {
	app, err := NewApp(ctx, conf)
	if err != nil {
		return err
	}
	// r.Use(app.Authorization.HandleAuthorization)
	wa := &WrapAuth{false, app.SessionAuthorizer.Authorize}
	sec := &s.SecurityConfig{
		SecuritySkip: conf.SecuritySkip,
		Check:        wa.Check,
		Authorize:    app.Authorizer.Authorize,
	}

	Handle(r, "/health", app.Health.Check, c.GET)
	Handle(r, "/authenticate", app.Authentication.Authenticate, c.POST)
	Handle(r, "/authentication/signout", app.Authentication.Logout, c.GET)

	r.Handle("/code/{code}", app.AuthorizationChecker.Check(http.HandlerFunc(app.Code.Load))).Methods(c.GET)
	r.Handle("/settings", app.AuthorizationChecker.Check(http.HandlerFunc(app.Settings.Save))).Methods(c.PATCH)

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

	m.HandleWithSecurity(users, "/search", app.User.Search, sec.Check, sec.Authorize, user, c.ActionRead, c.GET, c.POST)
	HandleWithSecurity(sec, users, "/{userId}", app.User.Load, user, c.ActionRead, c.GET)
	HandleWithSecurity(sec, users, "", app.User.Create, user, c.ActionWrite, c.POST)
	HandleWithSecurity(sec, users, "/{userId}", app.User.Update, user, c.ActionWrite, c.PUT)
	HandleWithSecurity(sec, users, "/{userId}", app.User.Patch, user, c.ActionWrite, c.PATCH)
	HandleWithSecurity(sec, users, "/{userId}", app.User.Delete, user, c.ActionWrite, c.DELETE)

	currencies := r.PathPrefix("/currencies").Subrouter()
	HandleWithSecurity(sec, currencies, "/search", app.Currency.Search, currency, c.ActionRead, c.GET, c.POST)
	HandleWithSecurity(sec, currencies, "/{currencyId}", app.Currency.Load, currency, c.ActionRead, c.GET)
	HandleWithSecurity(sec, currencies, "", app.Currency.Create, currency, c.ActionWrite, c.POST)
	HandleWithSecurity(sec, currencies, "/{currencyId}", app.Currency.Update, currency, c.ActionWrite, c.PUT)
	HandleWithSecurity(sec, currencies, "/{currencyId}", app.Currency.Patch, currency, c.ActionWrite, c.PATCH)
	HandleWithSecurity(sec, currencies, "/{currencyId}", app.Currency.Delete, currency, c.ActionWrite, c.DELETE)

	locales := r.PathPrefix("/locales").Subrouter()
	HandleWithSecurity(sec, locales, "/search", app.Locale.Search, locale, c.ActionRead, c.GET, c.POST)
	HandleWithSecurity(sec, locales, "/{localeId}", app.Locale.Load, locale, c.ActionRead, c.GET)
	HandleWithSecurity(sec, locales, "", app.Locale.Create, locale, c.ActionWrite, c.POST)
	HandleWithSecurity(sec, locales, "/{localeId}", app.Locale.Update, locale, c.ActionWrite, c.PUT)
	HandleWithSecurity(sec, locales, "/{localeId}", app.Locale.Patch, locale, c.ActionWrite, c.PATCH)
	HandleWithSecurity(sec, locales, "/{localeId}", app.Locale.Delete, locale, c.ActionWrite, c.DELETE)

	countries := r.PathPrefix("/countries").Subrouter()
	HandleWithSecurity(sec, countries, "/search", app.Country.Search, Country, c.ActionRead, c.GET, c.POST)
	HandleWithSecurity(sec, countries, "/{countryId}", app.Country.Load, Country, c.ActionRead, c.GET)
	HandleWithSecurity(sec, countries, "", app.Country.Create, Country, c.ActionWrite, c.POST)
	HandleWithSecurity(sec, countries, "/{countryId}", app.Country.Update, Country, c.ActionWrite, c.PUT)
	HandleWithSecurity(sec, countries, "/{countryId}", app.Country.Patch, Country, c.ActionWrite, c.PATCH)
	HandleWithSecurity(sec, countries, "/{countryId}", app.Country.Delete, Country, c.ActionWrite, c.DELETE)

	HandleWithSecurity(sec, r, "/audit-logs", app.AuditLog.Search, audit_log, c.ActionRead, c.GET, c.POST)
	HandleWithSecurity(sec, r, "/audit-logs/search", app.AuditLog.Search, audit_log, c.ActionRead, c.GET, c.POST)
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
type WrapAuth struct {
	isSkip    bool
	Authorize func(next http.Handler, skipRefreshTTL bool) http.Handler
}

func (h WrapAuth) Check(next http.Handler) http.Handler {
	return h.Authorize(next, h.isSkip)
}
