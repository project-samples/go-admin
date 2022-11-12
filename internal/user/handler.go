package user

import (
	"context"
	"net/http"
	"reflect"

	sv "github.com/core-go/core"
	"github.com/core-go/core/builder"
	"github.com/core-go/search"
)

type UserTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetUserByRole(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(
	find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error),
	userService UserRepository,
	conf sv.WriterConfig,
	logError func(context.Context, string, ...map[string]interface{}),
	generateId func(context.Context) (string, error),
	validate func(context.Context, interface{}) ([]sv.ErrorMessage, error),
	tracking builder.TrackingConfig,
	writeLog func(context.Context, string, string, bool, string) error) UserTransport {
	modelType := reflect.TypeOf(User{})
	searchModelType := reflect.TypeOf(UserFilter{})
	builder := builder.NewBuilderWithIdAndConfig(generateId, modelType, tracking)
	patchHandler, params := sv.CreatePatchAndParams(modelType, conf.Status, logError, userService.Patch, validate, builder.Patch, conf.Action, writeLog)
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, writeLog)
	return &UserHandler{service: userService, builder: builder, PatchHandler: patchHandler, SearchHandler: searchHandler, Params: params}
}

type UserHandler struct {
	service UserRepository
	builder sv.Builder
	*sv.PatchHandler
	*search.SearchHandler
	*sv.Params
}

func (h *UserHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Load(r.Context(), id)
		sv.RespondModel(w, r, result, err, h.Error, nil)
	}
}
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user User
	er1 := sv.Decode(w, r, &user, h.builder.Create)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &user)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Create) {
			result, er3 := h.service.Create(r.Context(), &user)
			sv.AfterCreated(w, r, &user, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Create)
		}
	}
}
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user User
	er1 := sv.DecodeAndCheckId(w, r, &user, h.Keys, h.Indexes, h.builder.Update)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &user)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Update) {
			result, er3 := h.service.Update(r.Context(), &user)
			sv.HandleResult(w, r, &user, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Update)
		}
	}
}
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Delete(r.Context(), id)
		sv.HandleDelete(w, r, result, err, h.Error, h.Log, h.Resource, h.Action.Delete)
	}
}
func (h *UserHandler) GetUserByRole(w http.ResponseWriter, r *http.Request) {
	roleId := r.URL.Query().Get("roleId")
	if len(roleId) > 0 {
		result, err := h.service.GetUserByRole(r.Context(), roleId)
		sv.RespondModel(w, r, result, err, h.Error, nil)
	}
}
