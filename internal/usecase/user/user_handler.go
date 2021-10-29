package user

import (
	"context"
	"encoding/json"
	"github.com/core-go/search"
	"net/http"
	"reflect"

	sv "github.com/core-go/service"
	"github.com/core-go/service/model-builder"
)


type UserHandler interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	service UserService
	*search.SearchHandler
	*sv.Params
	builder *builder.ModelBuilder
}

func NewUserHandler(
		find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error),
		userService UserService,
		conf sv.WriterConfig,
		logError func(context.Context, string),
		generateId func(context.Context) (string, error),
		validate func(context.Context, interface{}) ([]sv.ErrorMessage, error),
		tracking builder.TrackingConfig,
		writeLog func(context.Context, string, string, bool, string) error) UserHandler {
	modelType := reflect.TypeOf(User{})
	searchModelType := reflect.TypeOf(UserFilter{})
	builder := builder.NewDefaultModelBuilderByConfig(generateId, modelType, tracking)
	params := sv.CreateParams(modelType, conf.Status, logError, validate, conf.Action, writeLog)
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, writeLog)
	return &userHandler{ service: userService, SearchHandler: searchHandler, Params: params, builder: builder}
}

func (h *userHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Load(r.Context(), id)
		sv.RespondModel(w, r, result, err, h.Error, nil)
	}
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user User
	er1 := Decode(w, r, &user, h.builder.BuildToInsert)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &user)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Create) {
			result, er3 := h.service.Create(r.Context(), &user)
			sv.AfterCreated(w, r, &user, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Create)
		}
	}
}

func (h *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user User
	er1 := DecodeAndCheckId(w, r, &user, h.Keys, h.Indexes, h.builder.BuildToUpdate)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &user)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Update) {
			result, er3 := h.service.Update(r.Context(), &user)
			sv.HandleResult(w, r, &user, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Update)
		}
	}
}

func (h *userHandler) Patch(w http.ResponseWriter, r *http.Request) {
	var user User
	r, json, er1 := BuildMapAndCheckId(w, r, &user, h.Keys, h.Indexes, h.builder.BuildToPatch)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &user)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Patch) {
			result, er3 := h.service.Patch(r.Context(), json)
			sv.HandleResult(w, r, json, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Patch)
		}
	}
}
func (h *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Delete(r.Context(), id)
		sv.HandleDelete(w, r, result, err, h.Error, h.Log, h.Resource, h.Action.Delete)
	}
}
func Decode(w http.ResponseWriter, r *http.Request, obj interface{}, options...func(context.Context, interface{}) (interface{}, error)) error {
	er1 := json.NewDecoder(r.Body).Decode(obj)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return er1
	}
	if len(options) > 0 && options[0] != nil {
		_ , er2 := options[0](r.Context(), obj)
		if er2 != nil {
			http.Error(w, er2.Error(), http.StatusInternalServerError)
		}
		return er2
	}
	return nil
}
func DecodeAndCheckId(w http.ResponseWriter, r *http.Request, obj interface{}, keysJson []string, mapIndex map[string]int, options...func(context.Context, interface{}) (interface{}, error)) error {
	er1 := Decode(w, r, obj)
	if er1 != nil {
		return er1
	}
	return CheckId(w, r, obj, keysJson, mapIndex, options...)
}
func CheckId(w http.ResponseWriter, r *http.Request, body interface{}, keysJson []string, mapIndex map[string]int, options...func(context.Context, interface{}) (interface{}, error)) error {
	err := sv.MatchId(r, body, keysJson, mapIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	if len(options) > 0 && options[0] != nil {
		_ , er2 := options[0](r.Context(), body)
		if er2 != nil {
			http.Error(w, er2.Error(), http.StatusInternalServerError)
		}
		return er2
	}
	return nil
}
func BuildMapAndCheckId(w http.ResponseWriter, r *http.Request, obj interface{}, keysJson []string, mapIndex map[string]int, options...func(context.Context, interface{}) (interface{}, error)) (*http.Request, map[string]interface{}, error) {
	r2, body, er0 := sv.BuildFieldMapAndCheckId(w, r, obj, keysJson, mapIndex)
	if er0 != nil {
		return r2, body, er0
	}
	json, er1 := sv.BodyToJsonMap(r, obj, body, keysJson, mapIndex, options...)
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
	}
	return r2, json, er1
}
