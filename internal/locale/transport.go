package locale

import (
	"context"
	"database/sql"
	"net/http"
	"reflect"

	"github.com/core-go/core"
	v10 "github.com/core-go/core/v10"
	"github.com/core-go/search/query"
	q "github.com/core-go/sql"
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
	validator, err := v10.NewValidator()
	if err != nil {
		return nil, err
	}
	localeType := reflect.TypeOf(Locale{})
	queryLocale := query.UseQuery(db, "locale", localeType)
	localeSearchBuilder, err := q.NewSearchBuilder(db, localeType, queryLocale)
	if err != nil {
		return nil, err
	}
	localeRepository, err := q.NewRepository(db, "locale", localeType)
	if err != nil {
		return nil, err
	}
	localeService := NewLocaleService(localeRepository)
	localeHandler := NewLocaleHandler(localeSearchBuilder.Search, localeService, logError, validator.Validate, action)
	return localeHandler, nil
}
