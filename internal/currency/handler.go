package currency

import (
	"context"
	sv "github.com/core-go/core"
	s "github.com/core-go/search/handler"
)

func NewCurrencyHandler(find func(context.Context, *CurrencyFilter, int64, int64) ([]Currency, int64, error), service CurrencyService, logError sv.Log, validate func(context.Context, *Currency) ([]sv.ErrorMessage, error), writeLog sv.WriteLog, action *sv.ActionConfig) CurrencyTransport {
	hdl := sv.NewhandlerWithLog[Currency, string](service, logError, validate, action, writeLog)
	searchHandler := s.NewSearchHandler[Currency, *CurrencyFilter](find, logError, nil)
	return &CurrencyHandler{Handler: hdl, SearchHandler: searchHandler}
}

type CurrencyHandler struct {
	*sv.Handler[Currency, string]
	*s.SearchHandler[Currency, *CurrencyFilter]
}
