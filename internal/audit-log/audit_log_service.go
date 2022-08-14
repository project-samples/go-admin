package audit

import (
	"database/sql"
	"reflect"

	"github.com/core-go/search"
	"github.com/core-go/search/query"
	s "github.com/core-go/sql"
)

type AuditLogService interface {
	search.SearchService
}

type SqlAuditLogService struct {
	search.SearchService
}

func NewAuditLogService(db *sql.DB) (*SqlAuditLogService, error) {
	var model AuditLog
	tableName := "auditLog"
	modelType := reflect.TypeOf(model)
	builder := query.NewBuilder(db, tableName, modelType)
	searchService, err := s.NewSearcherWithQuery(db, modelType, builder.BuildQuery)
	if err != nil {
		return nil, err
	}
	return &SqlAuditLogService{SearchService: searchService}, nil
}
