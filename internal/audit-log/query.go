package audit

import (
	"database/sql"
	"reflect"

	"github.com/core-go/search"
	"github.com/core-go/search/query"
	s "github.com/core-go/sql"
)

type AuditLogQuery interface {
	search.SearchService
}

type SqlAuditLogQuery struct {
	search.SearchService
}

func NewAuditLogQuery(db *sql.DB) (*SqlAuditLogQuery, error) {
	var model AuditLog
	tableName := "auditLog"
	modelType := reflect.TypeOf(model)
	builder := query.NewBuilder(db, tableName, modelType)
	searchService, err := s.NewSearcherWithQuery(db, modelType, builder.BuildQuery)
	if err != nil {
		return nil, err
	}
	return &SqlAuditLogQuery{SearchService: searchService}, nil
}
