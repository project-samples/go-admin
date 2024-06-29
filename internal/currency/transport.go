package currency

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

type CurrencyTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewCurrencyTransport(db *sql.DB, logError func(context.Context, string, ...map[string]interface{}), action *core.ActionConfig) (CurrencyTransport, error) {
	validator, err := v10.NewValidator()
	if err != nil {
		return nil, err
	}
	currencyType := reflect.TypeOf(Currency{})
	queryCurrency := query.UseQuery(db, "currency", currencyType)
	currencySearchBuilder, err := q.NewSearchBuilder(db, currencyType, queryCurrency)
	if err != nil {
		return nil, err
	}
	currencyRepository, err := q.NewRepository(db, "currency", currencyType)
	if err != nil {
		return nil, err
	}
	currencyService := NewCurrencyService(currencyRepository)
	currencyHandler := NewCurrencyHandler(currencySearchBuilder.Search, currencyService, logError, validator.Validate, action)
	return currencyHandler, nil
}
