package locale

import (
	"github.com/core-go/core"
	s "github.com/core-go/search/handler"
)

func NewLocaleHandler(search s.Search[Locale, *LocaleFilter], service LocaleService, logError core.Log, validate core.Validate[*Locale], writeLog core.WriteLog, action *core.ActionConfig) LocaleTransport {
	hdl := core.NewhandlerWithLog[Locale, string](service, logError, validate, writeLog, action)
	searchHandler := s.NewSearchHandler[Locale, *LocaleFilter](search, logError, nil)
	return &LocaleHandler{Handler: hdl, SearchHandler: searchHandler}
}

type LocaleHandler struct {
	*core.Handler[Locale, string]
	*s.SearchHandler[Locale, *LocaleFilter]
}
