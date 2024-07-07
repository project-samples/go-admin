package locale

import (
	"context"
	sv "github.com/core-go/core"
	. "github.com/core-go/core/handler"
	"github.com/core-go/search/handler"
)

func NewLocaleHandler(find func(context.Context, *LocaleFilter, int64, int64) ([]Locale, int64, error), service LocaleService, logError sv.Log, validate func(context.Context, *Locale) ([]sv.ErrorMessage, error), writeLog sv.WriteLog, action *sv.ActionConf) LocaleTransport {
	hdl := NewhandlerWithLog[Locale, string](service, logError, validate, action, writeLog)
	searchHandler := search.NewSearchHandler[Locale, *LocaleFilter](find, logError, nil)
	return &LocaleHandler{Handler: hdl, SearchHandler: searchHandler}
}

type LocaleHandler struct {
	*Handler[Locale, string]
	*search.SearchHandler[Locale, *LocaleFilter]
}
