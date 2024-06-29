package country

import (
	"context"
	"net/http"
	"reflect"

	sv "github.com/core-go/core"
	"github.com/core-go/search"
)

func NewCountryHandler(find func(context.Context, interface{}, interface{}, int64, int64) (int64, error), service CountryService, logError sv.Log, validate func(context.Context, interface{}) ([]sv.ErrorMessage, error), action *sv.ActionConfig) CountryTransport {
	filterType := reflect.TypeOf(CountryFilter{})
	modelType := reflect.TypeOf(Country{})
	params := sv.CreateParams(modelType, logError, validate, action)
	searchHandler := search.NewSearchHandler(find, modelType, filterType, logError, params.Log)
	return &CountryHandler{service: service, SearchHandler: searchHandler, Params: params}
}

type CountryHandler struct {
	service CountryService
	*search.SearchHandler
	*sv.Params
}

func (h *CountryHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		res, err := h.service.Load(r.Context(), id)
		sv.Return(w, r, res, err, h.Error, nil)
	}
}
func (h *CountryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var locale Country
	er1 := sv.Decode(w, r, &locale)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &locale)
		if !sv.HasError(w, r, errors, er2, h.Error, h.Log, h.Resource, h.Action.Create) {
			res, er3 := h.service.Create(r.Context(), &locale)
			sv.AfterCreated(w, r, &locale, res, er3, h.Error, h.Log, h.Resource, h.Action.Create)
		}
	}
}
func (h *CountryHandler) Update(w http.ResponseWriter, r *http.Request) {
	var locale Country
	er1 := sv.DecodeAndCheckId(w, r, &locale, h.Keys, h.Indexes)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &locale)
		if !sv.HasError(w, r, errors, er2, h.Error, h.Log, h.Resource, h.Action.Update) {
			res, er3 := h.service.Update(r.Context(), &locale)
			sv.HandleResult(w, r, &locale, res, er3, h.Error, h.Log, h.Resource, h.Action.Update)
		}
	}
}
func (h *CountryHandler) Patch(w http.ResponseWriter, r *http.Request) {
	var locale Country
	r, json, er1 := sv.BuildMapAndCheckId(w, r, &locale, h.Keys, h.Indexes)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &locale)
		if !sv.HasError(w, r, errors, er2, h.Error, h.Log, h.Resource, h.Action.Patch) {
			res, er3 := h.service.Patch(r.Context(), json)
			sv.HandleResult(w, r, json, res, er3, h.Error, h.Log, h.Resource, h.Action.Patch)
		}
	}
}
func (h *CountryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		res, err := h.service.Delete(r.Context(), id)
		sv.HandleDelete(w, r, res, err, h.Error, h.Log, h.Resource, h.Action.Delete)
	}
}
