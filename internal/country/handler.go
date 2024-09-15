package country

import (
	"context"
	sv "github.com/core-go/core"
	s "github.com/core-go/search/handler"
)

func NewCountryHandler(find func(context.Context, *CountryFilter, int64, int64) ([]Country, int64, error), service CountryService, logError sv.Log, validate func(context.Context, *Country) ([]sv.ErrorMessage, error), writeLog sv.WriteLog, action *sv.ActionConfig) CountryTransport {
	hdl := sv.NewhandlerWithLog[Country, string](service, logError, validate, action, writeLog)
	searchHandler := s.NewSearchHandler[Country, *CountryFilter](find, logError, nil)
	return &CountryHandler{Handler: hdl, SearchHandler: searchHandler}
}

type CountryHandler struct {
	*sv.Handler[Country, string]
	*s.SearchHandler[Country, *CountryFilter]
}
