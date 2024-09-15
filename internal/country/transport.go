package country

import (
	"database/sql"
	"github.com/core-go/core"
	sv "github.com/core-go/core/sql"
	val "github.com/core-go/core/validator"
	"github.com/core-go/sql/adapter"
	"github.com/core-go/sql/query/builder"
	"net/http"
)

type CountryTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewCountryTransport(db *sql.DB, logError core.Log, writeLog core.WriteLog, action *core.ActionConfig) (CountryTransport, error) {
	validator, err := val.NewValidator[*Country]()
	if err != nil {
		return nil, err
	}
	queryCountry := builder.UseQuery[Country, *CountryFilter](db, "country")
	countrySearchBuilder, err := adapter.NewSearchAdapter[Country, string, *CountryFilter](db, "country", queryCountry)
	if err != nil {
		return nil, err
	}
	countryRepository, err := adapter.NewAdapter[Country, string](db, "country")
	if err != nil {
		return nil, err
	}
	countryService := sv.NewService[Country, string](db, countryRepository)
	countryHandler := NewCountryHandler(countrySearchBuilder.Search, countryService, logError, validator.Validate, writeLog, action)
	return countryHandler, nil
}
