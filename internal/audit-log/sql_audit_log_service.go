package audit

import (
	"database/sql"
	"reflect"

	"github.com/core-go/search"
	s "github.com/core-go/sql"
	"github.com/core-go/sql/query"
)

type SqlAuditLogService struct {
	search.SearchService
}

func NewAuditLogService(db *sql.DB) *SqlAuditLogService {
	var model AuditLog
	tableName := "auditLog"
	modelType := reflect.TypeOf(model)
	builder := query.NewBuilder(db, tableName, modelType)
	searchService := s.NewSearcherWithQuery(db, modelType, builder.BuildQuery)
	return &SqlAuditLogService{SearchService: searchService}
}
