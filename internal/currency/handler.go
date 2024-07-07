package currency

import (
	"context"
	sv "github.com/core-go/core"
	"github.com/core-go/core/handler"
	s "github.com/core-go/search/handler"
)

func NewCurrencyHandler(find func(context.Context, *CurrencyFilter, int64, int64) ([]Currency, int64, error), service CurrencyService, logError sv.Log, validate func(context.Context, *Currency) ([]sv.ErrorMessage, error), writeLog sv.WriteLog, action *sv.ActionConf) CurrencyTransport {
	hdl := handler.NewhandlerWithLog[Currency, string](service, logError, validate, action, writeLog)
	searchHandler := s.NewSearchHandler[Currency, *CurrencyFilter](find, logError, nil)
	return &CurrencyHandler{Handler: hdl, SearchHandler: searchHandler}
}

type CurrencyHandler struct {
	*handler.Handler[Currency, string]
	*s.SearchHandler[Currency, *CurrencyFilter]
}
