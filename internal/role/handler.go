package role

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/core-go/core"
	b "github.com/core-go/core/builder"
	search "github.com/core-go/search/handler"
)

func NewRoleHandler(
	find search.Search[Role, *RoleFilter],
	roleService RoleService,
	logError core.Log,
	validate core.Validate[*Role],
	tracking b.TrackingConfig,
	writeLog core.WriteLog,
	action *core.ActionConfig,
) *RoleHandler {
	roleType := reflect.TypeOf(Role{})
	builder := b.NewBuilderByConfig[Role](nil, tracking)
	params := core.CreateAttributes(roleType, logError, action, writeLog)
	searchHandler := search.NewSearchHandler[Role, *RoleFilter](find, logError, nil)
	return &RoleHandler{SearchHandler: searchHandler, service: roleService, validate: validate, builder: builder, Attributes: params}
}

type RoleHandler struct {
	service RoleRepository
	*search.SearchHandler[Role, *RoleFilter]
	*core.Attributes
	validate core.Validate[*Role]
	builder  core.Builder[Role]
}

func (h *RoleHandler) Load(w http.ResponseWriter, r *http.Request) {
	id, err := core.GetRequiredString(w, r)
	if err == nil {
		role, err := h.service.Load(r.Context(), id)
		if err != nil {
			h.Error(r.Context(), err.Error())
			http.Error(w, core.InternalServerError, http.StatusInternalServerError)
			return
		}
		if role == nil {
			core.JSON(w, http.StatusNotFound, role)
		} else {
			core.JSON(w, http.StatusOK, role)
		}
	}
}
func (h *RoleHandler) Create(w http.ResponseWriter, r *http.Request) {
	role, er1 := core.Decode[Role](w, r, h.builder.Create)
	if er1 == nil {
		errors, er2 := h.validate(r.Context(), &role)
		if !core.HasError(w, r, errors, er2, h.Error, &role, h.Log, h.Resource, h.Action.Create) {
			res, er3 := h.service.Create(r.Context(), &role)
			if er3 != nil {
				h.Error(r.Context(), er3.Error())
				h.Log(r.Context(), h.Resource, h.Action.Update, false, er3.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
				return
			}

			if res > 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, true, fmt.Sprintf("delete '%s'", role.RoleId))
				core.JSON(w, http.StatusCreated, role)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("conflict '%s'", role.RoleId))
				core.JSON(w, http.StatusConflict, res)
			}
		}
	}
}
func (h *RoleHandler) Update(w http.ResponseWriter, r *http.Request) {
	role, err := core.DecodeAndCheckId[Role](w, r, h.Keys, h.Indexes, h.builder.Update)
	if err == nil {
		errors, err := h.validate(r.Context(), &role)
		if !core.HasError(w, r, errors, err, h.Error, &role, h.Log, h.Resource, h.Action.Update) {
			res, err := h.service.Update(r.Context(), &role)
			if err != nil {
				h.Error(r.Context(), err.Error())
				h.Log(r.Context(), h.Resource, h.Action.Update, false, err.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
				return
			}

			if res > 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, true, fmt.Sprintf("%s '%s'", h.Action.Update, role.RoleId))
				core.JSON(w, http.StatusOK, role)
			} else if res == 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("not found '%s'", role.RoleId))
				core.JSON(w, http.StatusNotFound, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("conflict '%s'", role.RoleId))
				core.JSON(w, http.StatusConflict, res)
			}
		}
	}
}
func (h *RoleHandler) Patch(w http.ResponseWriter, r *http.Request) {
	r, role, jsonRole, err := core.BuildMapAndCheckId[Role](w, r, h.Keys, h.Indexes, h.builder.Update)
	if err == nil {
		errors, err := h.validate(r.Context(), &role)
		if !core.HasError(w, r, errors, err, h.Error, jsonRole, h.Log, h.Resource, h.Action.Patch) {
			res, err := h.service.Patch(r.Context(), jsonRole)
			if err != nil {
				h.Error(r.Context(), err.Error())
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, err.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
				return
			}

			if res > 0 {
				h.Log(r.Context(), h.Resource, h.Action.Patch, true, fmt.Sprintf("%s '%s'", h.Action.Patch, role.RoleId))
				core.JSON(w, http.StatusOK, jsonRole)
			} else if res == 0 {
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, fmt.Sprintf("not found '%s'", role.RoleId))
				core.JSON(w, http.StatusNotFound, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, fmt.Sprintf("conflict '%s'", role.RoleId))
				core.JSON(w, http.StatusConflict, res)
			}
		}
	}
}
func (h *RoleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := core.GetRequiredString(w, r)
	if err == nil {
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
func (h *RoleHandler) AssignRole(w http.ResponseWriter, r *http.Request) {
	id, err := core.GetRequiredString(w, r, 1)
	if err == nil {
		users, err := core.Decode[[]string](w, r)
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
