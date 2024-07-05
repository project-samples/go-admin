package user

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/core-go/core"
	b "github.com/core-go/core/builder"
	search "github.com/core-go/search/handler"
)

func NewUserHandler(
	find func(context.Context, *UserFilter, int64, int64) ([]User, int64, error),
	userService UserService,
	generateId core.Generate,
	validate core.Validate,
	logError core.Log,
	writeLog core.WriteLog,
	action *core.ActionConfig,
	tracking b.TrackingConfig,
) *UserHandler {
	userType := reflect.TypeOf(User{})
	builder := b.NewBuilderWithIdAndConfig(generateId, userType, tracking)
	patchHandler, params := core.CreatePatchAndParams(userType, logError, userService.Patch, validate, builder.Patch, action, writeLog)
	searchHandler := search.NewSearchHandler(find, logError, nil)
	return &UserHandler{service: userService, builder: builder, PatchHandler: patchHandler, SearchHandler: searchHandler, Params: params}
}

type UserHandler struct {
	service UserService
	builder core.Builder
	*core.PatchHandler
	*search.SearchHandler[User, *UserFilter]
	*core.Params
}

func (h *UserHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := core.GetRequiredParam(w, r)
	if len(id) > 0 {
		res, err := h.service.Load(r.Context(), id)
		if err != nil {
			h.Error(r.Context(), err.Error())
			http.Error(w, core.InternalServerError, http.StatusInternalServerError)
			return
		}
		if res == nil {
			core.JSON(w, http.StatusNotFound, res)
		} else {
			core.JSON(w, http.StatusOK, res)
		}
	}
}
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user User
	er1 := core.Decode(w, r, &user, h.builder.Create)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &user)
		if !core.HasError(w, r, errors, er2, h.Error, h.Log, h.Resource, h.Action.Create) {
			res, er3 := h.service.Create(r.Context(), &user)
			core.AfterCreated(w, r, &user, res, er3, h.Error, h.Log, h.Resource, h.Action.Create)
		}
	}
}
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user User
	err := core.DecodeAndCheckId(w, r, &user, h.Keys, h.Indexes, h.builder.Update)
	if err == nil {
		errors, err := h.Validate(r.Context(), &user)
		if !core.HasError(w, r, errors, err, h.Error, h.Log, h.Resource, h.Action.Update) {
			res, err := h.service.Update(r.Context(), &user)
			core.HandleResult(w, r, &user, res, err, h.Error, h.Log, h.Resource, h.Action.Update)
		}
	}
}
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := core.GetRequiredParam(w, r)
	if len(id) > 0 {
		res, err := h.service.Delete(r.Context(), id)
		if err != nil {
			h.Error(r.Context(), err.Error())
			h.Log(r.Context(), h.Resource, h.Action.Delete, false, err.Error())
			http.Error(w, core.InternalServerError, http.StatusInternalServerError)
		} else if res == 0 {
			h.Log(r.Context(), h.Resource, h.Action.Delete, false, fmt.Sprintf("not found '%s'", id))
			core.JSON(w, http.StatusNotFound, res)
		} else if res < 0 {
			h.Log(r.Context(), h.Resource, h.Action.Delete, false, fmt.Sprintf("conflict '%s'", id))
			core.JSON(w, http.StatusConflict, res)
		} else {
			h.Log(r.Context(), h.Resource, h.Action.Delete, true, fmt.Sprintf("delete '%s'", id))
			core.JSON(w, http.StatusOK, res)
		}
	}
}
func (h *UserHandler) GetUserByRole(w http.ResponseWriter, r *http.Request) {
	roleId := r.URL.Query().Get("roleId")
	if len(roleId) == 0 {
		http.Error(w, "roleId cannot be empty", http.StatusBadRequest)
		return
	}
	res, err := h.service.GetUserByRole(r.Context(), roleId)
	if err != nil {
		h.Error(r.Context(), err.Error())
		http.Error(w, core.InternalServerError, http.StatusInternalServerError)
		return
	}
	core.JSON(w, http.StatusOK, res)
}
