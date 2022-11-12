package audit

import (
	"context"
	"reflect"

	"github.com/core-go/search"
)

type AuditLogHandler struct {
	*search.SearchHandler
	service AuditLogQuery
}

func NewAuditLogHandler(service AuditLogQuery, logError func(context.Context, string, ...map[string]interface{}), writeLog func(context.Context, string, string, bool, string) error) *AuditLogHandler {
	searchModelType := reflect.TypeOf(AuditLogFilter{})
	modelType := reflect.TypeOf(AuditLog{})
	searchHandler := search.NewSearchHandler(service.Search, modelType, searchModelType, logError, writeLog)
	return &AuditLogHandler{SearchHandler: searchHandler, service: service}
}
