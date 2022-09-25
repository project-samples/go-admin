package fund

import (
	"context"
	"net/http"
	"reflect"

	"github.com/core-go/core"
	"github.com/core-go/search"
)

type FundHandler interface {
	core.Handler
}

func NewFundHandler(find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error),
	writeLog func(context.Context, string, string, bool, string) error,
	service FundService,
	logError func(context.Context, string, ...map[string]interface{}),
	validate func(context.Context, interface{}) ([]core.ErrorMessage, error),
	status core.StatusConfig, action core.ActionConfig) FundHandler {
	searchModelType := reflect.TypeOf(FundFilter{})
	modelType := reflect.TypeOf(Fund{})
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, writeLog)
	params := core.CreateParams(modelType, &status, logError, validate, &action)
	handler := core.NewHandler(service, modelType, nil, logError, validate)
	return &fundHandler{SearchHandler: searchHandler, GenericHandler: handler, fundService: service, Error: logError, Params: params}
}

type fundHandler struct {
	fundService FundService
	*core.GenericHandler
	*search.SearchHandler
	*core.Params
	Error func(context.Context, string, ...map[string]interface{})
}

func (f *fundHandler) Patch(w http.ResponseWriter, r *http.Request) {
	var fund Fund
	r, json, er1 := core.BuildMapAndCheckId(w, r, &fund, f.Keys, f.Params.Indexes)
	if er1 == nil {
		// errors, er2 := f.Validate(r.Context(), &fund)
		// fmt.Println(errors, er2)

		result, er3 := f.fundService.Patch(r.Context(), json)
		core.HandleResult(w, r, json, result, er3, f.Params.Status, f.Error, nil, f.Resource, f.Params.Action.Patch)
	}
}
