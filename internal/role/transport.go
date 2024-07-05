package role

import (
	"context"
	"database/sql"
	"reflect"

	"github.com/core-go/core"
	"github.com/core-go/core/builder"
	"github.com/core-go/core/shortid"
	"github.com/core-go/core/unique"
	v10 "github.com/core-go/core/v10"
	"github.com/core-go/search/convert"
	q "github.com/core-go/sql"
	"github.com/core-go/sql/query"
	"github.com/core-go/sql/template"
	tb "github.com/core-go/sql/template/builder"
)

func NewRoleTransport(db *sql.DB, checkDelete string, logError func(context.Context, string, ...map[string]interface{}), templates map[string]*template.Template, tracking builder.TrackingConfig, writeLog func(context.Context, string, string, bool, string) error, action *core.ActionConfig) (RoleTransport, error) {
	validator, err := v10.NewValidator()
	if err != nil {
		return nil, err
	}
	buildParam := q.GetBuild(db)
	roleType := reflect.TypeOf(Role{})
	queryRole, err := tb.UseQuery[*RoleFilter]("role", templates, &roleType, convert.ToMap, buildParam, q.GetSort)
	if err != nil {
		return nil, err
	}
	roleSearchBuilder, err := query.NewSearchBuilder[Role, *RoleFilter](db, queryRole)
	if err != nil {
		return nil, err
	}
	// roleValidator := user.NewRoleValidator(db, conf.Sql.Role.Duplicate, validator.validateFileName)
	roleValidator := unique.NewUniqueFieldValidator(db, "roles", "rolename", reflect.TypeOf(Role{}), validator.Validate)
	roleRepository, er6 := NewRoleAdapter(db, checkDelete) // cfg.Sql.Role.Check)
	if er6 != nil {
		return nil, er6
	}
	roleService := NewRoleService(roleRepository)
	roleHandler := NewRoleHandler(roleSearchBuilder.Search, roleService, shortid.Generate, roleValidator.Validate, logError, writeLog, action)
	return roleHandler, nil
}
