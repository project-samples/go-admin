package locale

import (
	sv "github.com/core-go/core"
	"github.com/core-go/search/handler"
)

func NewLocaleHandler(find search.Search[Locale, *LocaleFilter], service LocaleService, logError sv.Log, validate sv.Validate[*Locale], writeLog sv.WriteLog, action *sv.ActionConfig) LocaleTransport {
	hdl := sv.NewhandlerWithLog[Locale, string](service, logError, validate, action, writeLog)
	searchHandler := search.NewSearchHandler[Locale, *LocaleFilter](find, logError, nil)
	return &LocaleHandler{Handler: hdl, SearchHandler: searchHandler}
}

type LocaleHandler struct {
	*sv.Handler[Locale, string]
	*search.SearchHandler[Locale, *LocaleFilter]
}
