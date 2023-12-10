package audit

import (
	"net/http"
	"reflect"

	"github.com/core-go/core"
	s "github.com/core-go/search"
)

func NewAuditLogHandler(auditLogQuery AuditLogQuery, logError core.Log) *AuditLogHandler {
	paramIndex, filterIndex := s.BuildParams(reflect.TypeOf(AuditLogFilter{}))
	return &AuditLogHandler{
		query:       auditLogQuery,
		logError:    logError,
		paramIndex:  paramIndex,
		filterIndex: filterIndex,
	}
}

type AuditLogHandler struct {
	query       AuditLogQuery
	logError    core.Log
	paramIndex  map[string]int
	filterIndex int
}

func (h *AuditLogHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := core.GetRequiredParam(w, r)
	if len(id) > 0 {
		res, err := h.query.Load(r.Context(), id)
		if err == nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if res == nil {
			core.JSON(w, http.StatusNotFound, res)
		} else {
			core.JSON(w, http.StatusOK, res)
		}
	}
}

func (h *AuditLogHandler) Search(w http.ResponseWriter, r *http.Request) {
	filter := AuditLogFilter{Filter: &s.Filter{}}
	err := s.Decode(r, &filter, h.paramIndex, h.filterIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logs, total, err := h.query.Search(r.Context(), &filter)
	if err != nil {
		h.logError(r.Context(), err.Error())
		http.Error(w, core.InternalServerError, http.StatusInternalServerError)
		return
	}
	core.JSON(w, http.StatusOK, &s.Result{List: &logs, Total: total})
}
