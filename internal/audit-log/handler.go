package audit

import (
	"context"
	"net/http"
	"reflect"

	sv "github.com/core-go/core"
	s "github.com/core-go/search"
)

func NewAuditLogHandler(
	auditLogQuery AuditLogQuery,
	logError func(context.Context, string, ...map[string]interface{}),
) *AuditLogHandler {
	paramIndex, filterIndex := s.BuildParams(reflect.TypeOf(AuditLogFilter{}))
	return &AuditLogHandler{
		auditLogQuery: auditLogQuery,
		logError:      logError,
		paramIndex:    paramIndex,
		filterIndex:   filterIndex,
	}
}

type AuditLogHandler struct {
	auditLogQuery AuditLogQuery
	logError      func(context.Context, string, ...map[string]interface{})
	paramIndex    map[string]int
	filterIndex   int
}

func (h *AuditLogHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		res, err := h.auditLogQuery.Load(r.Context(), id)
		if err == nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if res == nil {
			sv.JSON(w, http.StatusNotFound, res)
			return
		}
		sv.JSON(w, http.StatusOK, res)
	}
}

func (h *AuditLogHandler) Search(w http.ResponseWriter, r *http.Request) {
	filter := AuditLogFilter{Filter: &s.Filter{}}
	err := s.Decode(r, &filter, h.paramIndex, h.filterIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	auditLogs, total, err := h.auditLogQuery.Search(r.Context(), &filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sv.JSON(w, http.StatusOK, &s.Result{List: &auditLogs, Total: total})
}
