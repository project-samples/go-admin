package currency

import (
	"github.com/core-go/core"
	s "github.com/core-go/search/handler"
)

func NewCurrencyHandler(search s.Search[Currency, *CurrencyFilter], service CurrencyService, logError core.Log, validate core.Validate[*Currency], writeLog core.WriteLog, action *core.ActionConfig) CurrencyTransport {
	handler := core.NewhandlerWithLog[Currency, string](service, logError, validate, writeLog, action)
	searchHandler := s.NewSearchHandler[Currency, *CurrencyFilter](search, logError, nil)
	return &CurrencyHandler{Handler: handler, SearchHandler: searchHandler}
}

type CurrencyHandler struct {
	*core.Handler[Currency, string]
	*s.SearchHandler[Currency, *CurrencyFilter]
}
