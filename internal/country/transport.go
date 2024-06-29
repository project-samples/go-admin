package country

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

type CountryTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewCountryTransport(db *sql.DB, logError func(context.Context, string, ...map[string]interface{}), action *core.ActionConfig) (CountryTransport, error) {
	validator, err := v10.NewValidator()
	if err != nil {
		return nil, err
	}
	countryType := reflect.TypeOf(Country{})
	queryCountry := query.UseQuery(db, "country", countryType)
	countrySearchBuilder, err := q.NewSearchBuilder(db, countryType, queryCountry)
	if err != nil {
		return nil, err
	}
	countryRepository, err := q.NewRepository(db, "country", countryType)
	if err != nil {
		return nil, err
	}
	countryService := NewCountryService(countryRepository)
	countryHandler := NewCountryHandler(countrySearchBuilder.Search, countryService, logError, validator.Validate, action)
	return countryHandler, nil
}
