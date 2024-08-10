package user

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/core-go/core"
	hdl "github.com/core-go/core/handler"
	b "github.com/core-go/core/handler/builder"
	v "github.com/core-go/core/validator"
	search "github.com/core-go/search/handler"
)

func NewUserHandler(
	find search.Search[User, *UserFilter],
	userService UserService,
	logError core.Log,
	validate v.Validate[*User],
	tracking b.TrackingConfig,
	writeLog core.WriteLog,
	action *core.ActionConfig,
) *UserHandler {
	userType := reflect.TypeOf(User{})
	builder := b.NewBuilderByConfig[User](nil, tracking)
	params := core.CreateParams(userType, logError, action, writeLog)
	searchHandler := search.NewSearchHandler[User, *UserFilter](find, logError, nil)
	return &UserHandler{SearchHandler: searchHandler, service: userService, validate: validate, builder: builder, Params: params}
}

type UserHandler struct {
	service UserRepository
	*search.SearchHandler[User, *UserFilter]
	*core.Params
	validate v.Validate[*User]
	builder  hdl.Builder[User]
}

func (h *UserHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := core.GetRequiredParam(w, r)
	if len(id) > 0 {
		user, err := h.service.Load(r.Context(), id)
		if err != nil {
			h.Error(r.Context(), err.Error())
			http.Error(w, core.InternalServerError, http.StatusInternalServerError)
			return
		}
		if user == nil {
			core.JSON(w, http.StatusNotFound, user)
		} else {
			core.JSON(w, http.StatusOK, user)
		}
	}
}
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	user, er1 := hdl.Decode(w, r, h.builder.Create)
	if er1 == nil {
		errors, er2 := h.validate(r.Context(), &user)
		if !core.HasError(w, r, errors, er2, h.Error, h.Log, h.Resource, h.Action.Create) {
			res, er3 := h.service.Create(r.Context(), &user)
			if er3 != nil {
				h.Error(r.Context(), er3.Error())
				h.Log(r.Context(), h.Resource, h.Action.Update, false, er3.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
				return
			}

			if res > 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, true, fmt.Sprintf("delete '%s'", user.UserId))
				core.JSON(w, http.StatusCreated, user)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("conflict '%s'", user.UserId))
				core.JSON(w, http.StatusConflict, res)
			}
		}
	}
}
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	user, err := hdl.DecodeAndCheckId[User](w, r, h.Keys, h.Indexes, h.builder.Update)
	if err == nil {
		errors, err := h.validate(r.Context(), &user)
		if !core.HasError(w, r, errors, err, h.Error, h.Log, h.Resource, h.Action.Update) {
			res, err := h.service.Update(r.Context(), &user)
			if err != nil {
				h.Error(r.Context(), err.Error())
				h.Log(r.Context(), h.Resource, h.Action.Update, false, err.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
				return
			}

			if res > 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, true, fmt.Sprintf("%s '%s'", h.Action.Update, user.UserId))
				core.JSON(w, http.StatusOK, user)
			} else if res == 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("not found '%s'", user.UserId))
				core.JSON(w, http.StatusNotFound, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("conflict '%s'", user.UserId))
				core.JSON(w, http.StatusConflict, res)
			}
		}
	}
}
func (h *UserHandler) Patch(w http.ResponseWriter, r *http.Request) {
	r, user, jsonUser, err := hdl.BuildMapAndCheckId[User](w, r, h.Keys, h.Indexes, h.builder.Update)
	if err == nil {
		errors, err := h.validate(r.Context(), &user)
		if !core.HasError(w, r, errors, err, h.Error, h.Log, h.Resource, h.Action.Patch) {
			res, err := h.service.Patch(r.Context(), jsonUser)
			if err != nil {
				h.Error(r.Context(), err.Error())
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, err.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
				return
			}

			if res > 0 {
				h.Log(r.Context(), h.Resource, h.Action.Patch, true, fmt.Sprintf("%s '%s'", h.Action.Patch, user.UserId))
				core.JSON(w, http.StatusOK, jsonUser)
			} else if res == 0 {
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, fmt.Sprintf("not found '%s'", user.UserId))
				core.JSON(w, http.StatusNotFound, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, fmt.Sprintf("conflict '%s'", user.UserId))
				core.JSON(w, http.StatusConflict, res)
			}
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
			return
		}

		if res > 0 {
			h.Log(r.Context(), h.Resource, h.Action.Delete, true, fmt.Sprintf("%s '%s'", h.Action.Delete, id))
			core.JSON(w, http.StatusOK, res)
		} else if res == 0 {
			h.Log(r.Context(), h.Resource, h.Action.Delete, false, fmt.Sprintf("not found '%s'", id))
			core.JSON(w, http.StatusNotFound, res)
		} else {
			h.Log(r.Context(), h.Resource, h.Action.Delete, false, fmt.Sprintf("conflict '%s'", id))
			core.JSON(w, http.StatusConflict, res)
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
