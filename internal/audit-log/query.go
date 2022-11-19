package audit

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/core-go/search"
	"github.com/core-go/search/convert"
	"github.com/core-go/search/template"
	q "github.com/core-go/sql"
)

type AuditLogQuery interface {
	Search(ctx context.Context, filter *AuditLogFilter) ([]AuditLog, int64, error)
	Load(ctx context.Context, id string) (*AuditLog, error)
}

type SqlAuditLogQuery struct {
	driver      string
	db          *sql.DB
	buildParam  func(int) string
	modelType   reflect.Type
	fieldsIndex map[string]int
	templates   map[string]*template.Template
}

func NewAuditLogQuery(
	db *sql.DB,
	templates map[string]*template.Template,
) (AuditLogQuery, error) {
	modelType := reflect.TypeOf(AuditLog{})
	driver := q.GetDriver(db)
	buildParam := q.GetBuild(db)
	fieldsIndex, err := q.GetColumnIndexes(modelType)
	if err != nil {
		return nil, err
	}
	return &SqlAuditLogQuery{
		db:          db,
		driver:      driver,
		buildParam:  buildParam,
		modelType:   modelType,
		fieldsIndex: fieldsIndex,
		templates:   templates,
	}, nil
}

func (s *SqlAuditLogQuery) Load(ctx context.Context, id string) (*AuditLog, error) {
	var rows []AuditLog
	query := fmt.Sprintf("select * from auditlog where id = %s limit 1", s.buildParam(1))
	err := q.Query(ctx, s.db, s.fieldsIndex, &rows, query, id)
	if err != nil {
		return nil, err
	}
	if len(rows) > 0 {
		return &rows[0], nil
	}
	return nil, nil
}
func (s SqlAuditLogQuery) Search(ctx context.Context, filter *AuditLogFilter) ([]AuditLog, int64, error) {
	var rows []AuditLog
	if filter.Limit <= 0 {
		return rows, 0, nil
	}
	ftr := convert.ToMap(filter, &s.modelType)

	query, params := template.Build(ftr, *s.templates["audit_log"], s.buildParam)
	offset := search.GetOffset(filter.Limit, filter.Page)
	if offset < 0 {
		offset = 0
	}
	pagingQuery := q.BuildPagingQuery(query, filter.Limit, offset, s.driver)
	countQuery, params := q.BuildCountQuery(query, params)

	total, err := q.Count(ctx, s.db, countQuery, params...)
	if total == 0 || err != nil {
		return rows, total, err
	}

	err = q.Query(ctx, s.db, s.fieldsIndex, &rows, pagingQuery, params...)
	return rows, total, err
}
