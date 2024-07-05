package locale

import (
	"context"
	"database/sql"
	"github.com/core-go/core"
	sv "github.com/core-go/core/sql"
	val "github.com/core-go/core/validator"
	"github.com/core-go/sql/adapter"
	"github.com/core-go/sql/query/builder"
	"net/http"
)

type LocaleTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewLocaleTransport(db *sql.DB, logError func(context.Context, string, ...map[string]interface{}), action *core.ActionConfig) (LocaleTransport, error) {
	validator, err := val.NewValidator[*Locale]()
	if err != nil {
		return nil, err
	}
	queryLocale := builder.UseQuery[Locale, *LocaleFilter](db, "locale")
	localeSearchBuilder, err := adapter.NewSearchAdapter[Locale, string, *LocaleFilter](db, "locale", queryLocale)
	if err != nil {
		return nil, err
	}
	localeRepository, err := adapter.NewAdapter[Locale, string](db, "locale")
	if err != nil {
		return nil, err
	}
	localeService := sv.NewService[Locale, string](db, localeRepository)
	localeHandler := NewLocaleHandler(localeSearchBuilder.Search, localeService, logError, validator.Validate, action)
	return localeHandler, nil
}
