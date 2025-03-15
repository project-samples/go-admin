package country

import (
	"github.com/core-go/core"
	s "github.com/core-go/search/handler"
)

func NewCountryHandler(search s.Search[Country, *CountryFilter], service CountryService, logError core.Log, validate core.Validate[*Country], writeLog core.WriteLog, action *core.ActionConfig) CountryTransport {
	handler := core.NewhandlerWithLog[Country, string](service, logError, validate, writeLog, action)
	searchHandler := s.NewSearchHandler[Country, *CountryFilter](search, logError, nil)
	return &CountryHandler{Handler: handler, SearchHandler: searchHandler}
}

type CountryHandler struct {
	*core.Handler[Country, string]
	*s.SearchHandler[Country, *CountryFilter]
}
