package role

import (
	"context"
	"net/http"
	"reflect"

	"github.com/core-go/search"
	sv "github.com/core-go/service"
	"github.com/core-go/service/builder"
)

type RoleHandler interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewRoleHandler(
	find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error),
	roleService RoleService,
	conf sv.WriterConfig,
	logError func(context.Context, string),
	generateId func(context.Context) (string, error),
	validate func(context.Context, interface{}) ([]sv.ErrorMessage, error),
	tracking builder.TrackingConfig,
	writeLog func(context.Context, string, string, bool, string) error) RoleHandler {
	modelType := reflect.TypeOf(Role{})
	searchModelType := reflect.TypeOf(RoleFilter{})
	builder := builder.NewBuilderWithIdAndConfig(generateId, modelType, tracking)
	params := sv.CreateParams(modelType, conf.Status, logError, validate, conf.Action, writeLog)
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, writeLog)
	return &roleHandler{service: roleService, builder: builder, SearchHandler: searchHandler, Params: params}
}

type roleHandler struct {
	service RoleService
	builder sv.Builder
	*search.SearchHandler
	*sv.Params
}

func (h *roleHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Load(r.Context(), id)
		sv.RespondModel(w, r, result, err, h.Error, nil)
	}
}
func (h *roleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var role Role
	er1 := sv.Decode(w, r, &role, h.builder.Create)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &role)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Create) {
			result, er3 := h.service.Create(r.Context(), &role)
			sv.AfterCreated(w, r, &role, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Create)
		}
	}
}
func (h *roleHandler) Update(w http.ResponseWriter, r *http.Request) {
	var role Role
	er1 := sv.DecodeAndCheckId(w, r, &role, h.Keys, h.Indexes, h.builder.Update)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &role)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Update) {
			result, er3 := h.service.Update(r.Context(), &role)
			sv.HandleResult(w, r, &role, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Update)
		}
	}
}
func (h *roleHandler) Patch(w http.ResponseWriter, r *http.Request) {
	var role Role
	r, json, er1 := sv.BuildMapAndCheckId(w, r, &role, h.Keys, h.Indexes, h.builder.Patch)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &role)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Patch) {
			result, er3 := h.service.Patch(r.Context(), json)
			sv.HandleResult(w, r, json, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Patch)
		}
	}
}
func (h *roleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Delete(r.Context(), id)
		sv.HandleDelete(w, r, result, err, h.Error, h.Log, h.Resource, h.Action.Delete)
	}
}
