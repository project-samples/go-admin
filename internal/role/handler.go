package role

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/core-go/core"
	"github.com/core-go/core/builder"
	"github.com/core-go/search"
)

type RoleTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	AssignRole(w http.ResponseWriter, r *http.Request)
}

func NewRoleHandler(
	find core.Search,
	roleService RoleService,
	generateId core.Generate,
	validate core.Validate,
	logError core.Log,
	writeLog core.WriteLog,
	action *core.ActionConfig,
	tracking builder.TrackingConfig,
) *RoleHandler {
	roleType := reflect.TypeOf(Role{})
	searchModelType := reflect.TypeOf(RoleFilter{})
	builder := builder.NewBuilderWithIdAndConfig(generateId, roleType, tracking)
	params := core.CreateParams(roleType, logError, validate, action, writeLog)
	searchHandler := search.NewSearchHandler(find, roleType, searchModelType, logError, nil)
	return &RoleHandler{service: roleService, builder: builder, SearchHandler: searchHandler, Params: params}
}

type RoleHandler struct {
	service RoleRepository
	builder core.Builder
	*search.SearchHandler
	*core.Params
}

func (h *RoleHandler) Load(w http.ResponseWriter, r *http.Request) {
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
func (h *RoleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var role Role
	er1 := core.Decode(w, r, &role, h.builder.Create)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &role)
		if !core.HasError(w, r, errors, er2, h.Error, h.Log, h.Resource, h.Action.Create) {
			res, er3 := h.service.Create(r.Context(), &role)
			if er3 != nil {
				h.Error(r.Context(), er3.Error())
				h.Log(r.Context(), h.Resource, h.Action.Update, false, er3.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
			} else if res <= 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("conflict '%s'", role.RoleId))
				core.JSON(w, http.StatusConflict, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Update, true, fmt.Sprintf("delete '%s'", role.RoleId))
				core.JSON(w, http.StatusCreated, res)
			}
		}
	}
}
func (h *RoleHandler) Update(w http.ResponseWriter, r *http.Request) {
	var role Role
	err := core.DecodeAndCheckId(w, r, &role, h.Keys, h.Indexes, h.builder.Update)
	if err == nil {
		errors, err := h.Validate(r.Context(), &role)
		if !core.HasError(w, r, errors, err, h.Error, h.Log, h.Resource, h.Action.Update) {
			res, err := h.service.Update(r.Context(), &role)
			if err != nil {
				h.Error(r.Context(), err.Error())
				h.Log(r.Context(), h.Resource, h.Action.Update, false, err.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
			} else if res == 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("not found '%s'", role.RoleId))
				core.JSON(w, http.StatusNotFound, res)
			} else if res < 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("conflict '%s'", role.RoleId))
				core.JSON(w, http.StatusConflict, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Update, true, fmt.Sprintf("%s '%s'", h.Action.Update, role.RoleId))
				core.JSON(w, http.StatusOK, res)
			}
		}
	}
}
func (h *RoleHandler) Patch(w http.ResponseWriter, r *http.Request) {
	var role Role
	r, json, err := core.BuildMapAndCheckId(w, r, &role, h.Keys, h.Indexes, h.builder.Patch)
	if err == nil {
		errors, err := h.Validate(r.Context(), &role)
		if !core.HasError(w, r, errors, err, h.Error, h.Log, h.Resource, h.Action.Patch) {
			res, err := h.service.Patch(r.Context(), json)
			if err != nil {
				h.Error(r.Context(), err.Error())
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, err.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
			} else if res == 0 {
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, fmt.Sprintf("not found '%s'", role.RoleId))
				core.JSON(w, http.StatusNotFound, res)
			} else if res < 0 {
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, fmt.Sprintf("conflict '%s'", role.RoleId))
				core.JSON(w, http.StatusConflict, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Patch, true, fmt.Sprintf("%s '%s'", h.Action.Patch, role.RoleId))
				core.JSON(w, http.StatusOK, res)
			}
		}
	}
}
func (h *RoleHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
			h.Log(r.Context(), h.Resource, h.Action.Delete, true, fmt.Sprintf("%s '%s'", h.Action.Delete, id))
			core.JSON(w, http.StatusOK, res)
		}
	}
}

func (h *RoleHandler) AssignRole(w http.ResponseWriter, r *http.Request) {
	var users []string
	id := core.GetRequiredParam(w, r, 1)
	if len(id) > 0 {
		err := core.Decode(w, r, &users)
		if err == nil {
			res, err := h.service.AssignRole(r.Context(), id, users)
			if err != nil {
				h.Error(r.Context(), err.Error())
				h.Log(r.Context(), h.Resource, "assign", false, err.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
			} else if res == 0 {
				h.Log(r.Context(), h.Resource, "assign", false, fmt.Sprintf("not found '%s'", id))
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
			} else {
				h.Log(r.Context(), h.Resource, "assign", true, fmt.Sprintf("assign '%s'", id))
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
			}
		}
	}
}
