package locale

import (
	"context"
	sv "github.com/core-go/core"
	. "github.com/core-go/core/handler"
	search "github.com/core-go/search/handler"
)

func NewLocaleHandler(find func(context.Context, *LocaleFilter, int64, int64) ([]Locale, int64, error), service LocaleService, logError sv.Log, validate func(context.Context, *Locale) ([]sv.ErrorMessage, error), action *sv.ActionConfig) LocaleTransport {
	hdl := Newhandler[Locale, string](service, logError, validate)
	searchHandler := search.NewSearchHandler[Locale, *LocaleFilter](find, logError, nil)
	return &LocaleHandler{Handler: hdl, SearchHandler: searchHandler}
}

type LocaleHandler struct {
	*Handler[Locale, string]
	*search.SearchHandler[Locale, *LocaleFilter]
}
