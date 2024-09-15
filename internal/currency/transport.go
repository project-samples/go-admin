package currency

import (
	"database/sql"
	"github.com/core-go/core"
	sv "github.com/core-go/core/sql"
	val "github.com/core-go/core/validator"
	"github.com/core-go/sql/adapter"
	"github.com/core-go/sql/query/builder"
	"net/http"
)

type CurrencyTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewCurrencyTransport(db *sql.DB, logError core.Log, writeLog core.WriteLog, action *core.ActionConfig) (CurrencyTransport, error) {
	validator, err := val.NewValidator[*Currency]()
	if err != nil {
		return nil, err
	}
	queryCurrency := builder.UseQuery[Currency, *CurrencyFilter](db, "currency")
	currencySearchBuilder, err := adapter.NewSearchAdapter[Currency, string, *CurrencyFilter](db, "currency", queryCurrency)
	if err != nil {
		return nil, err
	}
	currencyRepository, err := adapter.NewAdapter[Currency, string](db, "currency")
	if err != nil {
		return nil, err
	}
	currencyService := sv.NewService[Currency, string](db, currencyRepository)
	currencyHandler := NewCurrencyHandler(currencySearchBuilder.Search, currencyService, logError, validator.Validate, writeLog, action)
	return currencyHandler, nil
}
