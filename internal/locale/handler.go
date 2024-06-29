package locale

import (
	"context"
	"net/http"
	"reflect"

	sv "github.com/core-go/core"
	"github.com/core-go/search"
)

func NewLocaleHandler(find func(context.Context, interface{}, interface{}, int64, int64) (int64, error), service LocaleService, logError sv.Log, validate func(context.Context, interface{}) ([]sv.ErrorMessage, error), action *sv.ActionConfig) LocaleTransport {
	filterType := reflect.TypeOf(LocaleFilter{})
	modelType := reflect.TypeOf(Locale{})
	params := sv.CreateParams(modelType, logError, validate, action)
	searchHandler := search.NewSearchHandler(find, modelType, filterType, logError, params.Log)
	return &LocaleHandler{service: service, SearchHandler: searchHandler, Params: params}
}

type LocaleHandler struct {
	service LocaleService
	*search.SearchHandler
	*sv.Params
}

func (h *LocaleHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		res, err := h.service.Load(r.Context(), id)
		sv.Return(w, r, res, err, h.Error, nil)
	}
}
func (h *LocaleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var locale Locale
	er1 := sv.Decode(w, r, &locale)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &locale)
		if !sv.HasError(w, r, errors, er2, h.Error, h.Log, h.Resource, h.Action.Create) {
			res, er3 := h.service.Create(r.Context(), &locale)
			sv.AfterCreated(w, r, &locale, res, er3, h.Error, h.Log, h.Resource, h.Action.Create)
		}
	}
}
func (h *LocaleHandler) Update(w http.ResponseWriter, r *http.Request) {
	var locale Locale
	er1 := sv.DecodeAndCheckId(w, r, &locale, h.Keys, h.Indexes)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &locale)
		if !sv.HasError(w, r, errors, er2, h.Error, h.Log, h.Resource, h.Action.Update) {
			res, er3 := h.service.Update(r.Context(), &locale)
			sv.HandleResult(w, r, &locale, res, er3, h.Error, h.Log, h.Resource, h.Action.Update)
		}
	}
}
func (h *LocaleHandler) Patch(w http.ResponseWriter, r *http.Request) {
	var locale Locale
	r, json, er1 := sv.BuildMapAndCheckId(w, r, &locale, h.Keys, h.Indexes)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &locale)
		if !sv.HasError(w, r, errors, er2, h.Error, h.Log, h.Resource, h.Action.Patch) {
			res, er3 := h.service.Patch(r.Context(), json)
			sv.HandleResult(w, r, json, res, er3, h.Error, h.Log, h.Resource, h.Action.Patch)
		}
	}
}
func (h *LocaleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		res, err := h.service.Delete(r.Context(), id)
		sv.HandleDelete(w, r, res, err, h.Error, h.Log, h.Resource, h.Action.Delete)
	}
}
