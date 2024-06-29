package currency

import (
	"context"
	"net/http"
	"reflect"

	"github.com/core-go/core"
	"github.com/core-go/search"
)

func NewCurrencyHandler(find func(context.Context, interface{}, interface{}, int64, int64) (int64, error), service CurrencyService, logError core.Log, validate func(context.Context, interface{}) ([]core.ErrorMessage, error), action *core.ActionConfig) CurrencyTransport {
	filterType := reflect.TypeOf(CurrencyFilter{})
	modelType := reflect.TypeOf(Currency{})
	params := core.CreateParams(modelType, logError, validate, action)
	searchHandler := search.NewSearchHandler(find, modelType, filterType, logError, params.Log)
	return &CurrencyHandler{service: service, SearchHandler: searchHandler, Params: params}
}

type CurrencyHandler struct {
	service CurrencyService
	*search.SearchHandler
	*core.Params
}

func (h *CurrencyHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := core.GetRequiredParam(w, r)
	if len(id) > 0 {
		res, err := h.service.Load(r.Context(), id)
		core.Return(w, r, res, err, h.Error, nil)
	}
}
func (h *CurrencyHandler) Create(w http.ResponseWriter, r *http.Request) {
	var currency Currency
	er1 := core.Decode(w, r, &currency)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &currency)
		if !core.HasError(w, r, errors, er2, h.Error, h.Log, h.Resource, h.Action.Create) {
			res, er3 := h.service.Create(r.Context(), &currency)
			core.AfterCreated(w, r, &currency, res, er3, h.Error, h.Log, h.Resource, h.Action.Create)
		}
	}
}
func (h *CurrencyHandler) Update(w http.ResponseWriter, r *http.Request) {
	var currency Currency
	er1 := core.DecodeAndCheckId(w, r, &currency, h.Keys, h.Indexes)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &currency)
		if !core.HasError(w, r, errors, er2, h.Error, h.Log, h.Resource, h.Action.Update) {
			res, er3 := h.service.Update(r.Context(), &currency)
			core.HandleResult(w, r, &currency, res, er3, h.Error, h.Log, h.Resource, h.Action.Update)
		}
	}
}
func (h *CurrencyHandler) Patch(w http.ResponseWriter, r *http.Request) {
	var currency Currency
	r, json, er1 := core.BuildMapAndCheckId(w, r, &currency, h.Keys, h.Indexes)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &currency)
		if !core.HasError(w, r, errors, er2, h.Error, h.Log, h.Resource, h.Action.Patch) {
			res, er3 := h.service.Patch(r.Context(), json)
			core.HandleResult(w, r, json, res, er3, h.Error, h.Log, h.Resource, h.Action.Patch)
		}
	}
}
func (h *CurrencyHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := core.GetRequiredParam(w, r)
	if len(id) > 0 {
		res, err := h.service.Delete(r.Context(), id)
		core.HandleDelete(w, r, res, err, h.Error, h.Log, h.Resource, h.Action.Delete)
	}
}
