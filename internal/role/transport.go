package role

import (
	"database/sql"
	"net/http"
	"reflect"

	"github.com/core-go/core"
	"github.com/core-go/core/handler/builder"
	"github.com/core-go/core/unique"
	v "github.com/core-go/core/validator"
	"github.com/core-go/search/convert"
	q "github.com/core-go/sql"
	"github.com/core-go/sql/query"
	"github.com/core-go/sql/template"
	tb "github.com/core-go/sql/template/builder"
)

type RoleTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	AssignRole(w http.ResponseWriter, r *http.Request)
}

func NewRoleTransport(db *sql.DB, checkDelete string, logError core.Log, templates map[string]*template.Template, tracking builder.TrackingConfig, writeLog core.WriteLog, action *core.ActionConfig) (RoleTransport, error) {
	validator, err := v.NewValidator[*Role]()
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
	roleValidator, err := unique.NewUniqueFieldValidator[*Role](db, "roles", "rolename", validator.Validate)
	if err != nil {
		return nil, err
	}
	roleRepository, er6 := NewRoleAdapter(db, checkDelete) // cfg.Sql.Role.Check)
	if er6 != nil {
		return nil, er6
	}
	roleService := NewRoleService(roleRepository)
	roleHandler := NewRoleHandler(roleSearchBuilder.Search, roleService, logError, roleValidator.Validate, tracking, writeLog, action)
	return roleHandler, nil
}
