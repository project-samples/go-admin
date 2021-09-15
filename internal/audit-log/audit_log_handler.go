package audit

import (
	"context"
	"reflect"

	"github.com/core-go/search"
)

type AuditLogHandler struct {
	*search.SearchHandler
	service AuditLogService
}

func NewAuditLogHandler(service AuditLogService, logError func(context.Context, string), writeLog func(context.Context, string, string, bool, string) error) *AuditLogHandler {
	searchModelType := reflect.TypeOf(AuditLogSM{})
	modelType := reflect.TypeOf(AuditLog{})
	searchHandler := search.NewSearchHandler(service.Search, modelType, searchModelType, logError, writeLog)
	return &AuditLogHandler{SearchHandler: searchHandler, service: service}
}
