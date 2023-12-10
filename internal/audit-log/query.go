package audit

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/core-go/core/user"
	"github.com/core-go/search"
	"github.com/core-go/search/convert"
	q "github.com/core-go/sql"
	"github.com/core-go/sql/template"
)

type AuditLogQuery interface {
	Search(ctx context.Context, filter *AuditLogFilter) ([]AuditLog, int64, error)
	Load(ctx context.Context, id string) (*AuditLog, error)
}

type SqlAuditLogQuery struct {
	driver       string
	db           *sql.DB
	buildParam   func(int) string
	AuditLogType reflect.Type
	Map          map[string]int
	Fields       string
	templates    map[string]*template.Template
	GetUsers     user.GetUsers
}

func NewAuditLogQuery(db *sql.DB, templates map[string]*template.Template, getUsers user.GetUsers) (AuditLogQuery, error) {
	logType := reflect.TypeOf(AuditLog{})
	fieldsIndex, fields, buildParam, driver, err := q.InitFields(logType, db)
	return &SqlAuditLogQuery{
		db:           db,
		driver:       driver,
		buildParam:   buildParam,
		AuditLogType: logType,
		Map:          fieldsIndex,
		Fields:       fields,
		templates:    templates,
		GetUsers:     getUsers,
	}, err
}

func (s *SqlAuditLogQuery) Load(ctx context.Context, id string) (*AuditLog, error) {
	var rows []AuditLog
	query := fmt.Sprintf("select %s from auditlog where id = %s limit 1", s.Fields, s.buildParam(1))
	err := q.Query(ctx, s.db, s.Map, &rows, query, id)
	if len(rows) > 0 {
		return &rows[0], err
	}
	return nil, err
}
func (s SqlAuditLogQuery) Search(ctx context.Context, filter *AuditLogFilter) ([]AuditLog, int64, error) {
	var rows []AuditLog
	if filter.Limit <= 0 {
		return rows, 0, nil
	}
	ftr := convert.ToMap(filter, &s.AuditLogType)
	ftr["fields"] = s.Fields

	query, params := template.Build(ftr, *s.templates["audit_log"], s.buildParam)
	offset := search.GetOffset(filter.Limit, filter.Page)
	pagingQuery := q.BuildPagingQuery(query, filter.Limit, offset, s.driver)
	countQuery := q.BuildCountQuery(query)

	total, err := q.Count(ctx, s.db, countQuery, params...)
	if total == 0 || err != nil {
		return rows, total, err
	}

	err = q.Query(ctx, s.db, s.Map, &rows, pagingQuery, params...)
	if len(rows) == 0 {
		return rows, total, err
	}
	if s.GetUsers != nil {
		ids := user.Unique(toListIds(rows))
		users, err := s.GetUsers(ctx, ids)
		if err != nil {
			return rows, total, err
		}
		usersMap := user.ToMap(users)
		for i, row := range rows {
			if u, ok := usersMap[row.UserId]; ok {
				rows[i].Email = u.Email
			}
		}
	}
	return rows, total, err
}

func toListIds(rows []AuditLog) []string {
	rs := make([]string, len(rows))
	for i, row := range rows {
		rs[i] = row.UserId
	}
	return rs
}
