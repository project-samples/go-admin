package role

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	sv "github.com/core-go/service"
	q "github.com/core-go/sql"
)

const ActionNone int32 = 0

type RoleService interface {
	Load(ctx context.Context, id string) (*Role, error)
	Create(ctx context.Context, role *Role) (int64, error)
	Update(ctx context.Context, role *Role) (int64, error)
	Patch(ctx context.Context, obj map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type roleService struct {
	db				 *sql.DB
	repository 		 sv.Repository
	BuildParam       func(int) string
	CheckDelete      string
	modelType        reflect.Type
	Driver           string
	Map              map[string]int
	roleSchema       *q.Schema
	roleModuleSchema *q.Schema
}

func NewRoleService(db *sql.DB, checkDelete string) (*roleService, error) {
	modelType := reflect.TypeOf(Role{})
	buildParam := q.GetBuild(db)
	repository,er1 := q.NewRepository(db, "roles", modelType)
	if er1 != nil {
		return nil, er1
	}
	var r RoleModule
	subType := reflect.TypeOf(r)
	m, err := q.GetColumnIndexes(subType)
	if err != nil {
		return nil, err
	}
	sql := q.ReplaceQueryArgs(q.GetDriver(db), checkDelete)
	roleSchema := q.CreateSchema(modelType)
	roleModuleSchema := q.CreateSchema(subType)
	driver := q.GetDriver(db)
	return &roleService{db: db, Driver: driver, repository: repository, BuildParam: buildParam, CheckDelete: sql, modelType: modelType, Map: m, roleSchema: roleSchema, roleModuleSchema: roleModuleSchema}, nil
}

func (s *roleService) Load(ctx context.Context, roleId string) (*Role, error) {
	var role Role
	ok, err := s.repository.LoadAndDecode(ctx, roleId, &role)
	if !ok || err != nil {
		return nil, err
	}

	privileges, er3 := getPrivileges(ctx, s.db, roleId, s.BuildParam, getModules, s.Map)
	if er3 != nil {
		return nil, er3
	}
	role.Privileges = privileges
	return &role, nil
}
func (s *roleService) Create(ctx context.Context, role *Role) (int64, error) {
	sts, err := buildInsertStatements(role, s.Driver, s.BuildParam, s.roleSchema, s.roleModuleSchema)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
}
func (s *roleService) Update(ctx context.Context, role *Role) (int64, error) {
	sts, err := buildUpdateStatements(role, s.Driver, s.BuildParam, s.roleSchema, s.roleModuleSchema)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
}
func (s *roleService) Delete(ctx context.Context, id string) (int64, error) {
	if len(s.CheckDelete) > 0 {
		exist, er0 := checkExist(s.db, s.CheckDelete, id)
		if exist || er0 != nil {
			return -1, er0
		}
	}
	sts, er1 := buildDeleteStatements(id, s.BuildParam)
	if er1 != nil {
		return 0, er1
	}
	return sts.Exec(ctx, s.db)
}
func (s *roleService) Patch(ctx context.Context, obj map[string]interface{}) (int64, error) {
	sts, err := buildPatchRoleStatements(obj, s.BuildParam, s.modelType)
	if err != nil {
		return 0, err
	}

	return sts.Exec(ctx, s.db)
}
func checkExist(db *sql.DB, sql string, args ...interface{}) (bool, error) {
	rows, err := db.Query(sql, args...)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		return true, nil
	}
	return false, nil
}
func buildInsertStatements(role *Role, driver string, buildParam func(int) string, roleSchema *q.Schema, roleModuleSchema *q.Schema) (q.Statements, error) {
	modules, er1 := buildModules(role.RoleId, role.Privileges)
	if er1 != nil {
		return nil, er1
	}
	sts := q.NewStatements(true)
	sts.Add(q.BuildToInsert("roles", role, buildParam, roleSchema))
	if modules != nil {
		query, args, er2 := q.BuildToInsertBatch("roleModules", modules, driver, roleModuleSchema)
		if er2 != nil {
			return nil, er2
		}
		sts.Add(query, args)
	}
	return sts, nil
}
func buildUpdateStatements(role *Role, driver string, buildParam func(int) string, roleSchema *q.Schema, roleModuleSchema *q.Schema) (q.Statements, error) {
	modules, err := buildModules(role.RoleId, role.Privileges)
	if err != nil {
		return nil, err
	}
	sts := q.NewStatements(true)
	sts.Add(q.BuildToUpdate("roles", role, buildParam, roleSchema))

	deleteModules := fmt.Sprintf("delete from roleModules where roleId = %s", buildParam(1))
	sts.Add(deleteModules, []interface{}{role.RoleId})
	if modules != nil {
		query, args, er2 := q.BuildToInsertBatch("roleModules", modules, driver, roleModuleSchema)
		if er2 != nil {
			return nil, er2
		}
		sts.Add(query, args)
	}
	return sts, nil
}
func buildDeleteStatements(roleId string, buildParam func(int) string) (q.Statements, error) {
	sts := q.NewStatements(false)

	deleteModules := fmt.Sprintf("delete from roleModules where roleId = %s", buildParam(1))
	sts.Add(deleteModules, []interface{}{roleId})

	deleteRole := fmt.Sprintf("delete from roles where roleId = %s", buildParam(1))
	sts.Add(deleteRole, []interface{}{roleId})

	return sts, nil
}
func buildModules(roleId string, privileges []string) ([]RoleModule, error) {
	if privileges == nil || len(privileges) <= 0 {
		return nil, nil
	}
	modules := make([]RoleModule, 0)
	for _, p := range privileges {
		m := toModules(p)
		m.RoleId = roleId
		modules = append(modules, m)
	}
	return modules, nil
}
func getPrivileges(ctx context.Context, db *sql.DB, roleId string, buildParam func(int) string, getModules func(context.Context, *sql.DB, string, func(int) string, map[string]int) ([]RoleModule, error), m map[string]int) ([]string, error) {
	modules, er1 := getModules(ctx, db, roleId, buildParam, m)
	if er1 != nil {
		return nil, er1
	}
	privileges := buildPrivileges(modules)
	return privileges, nil
}
func buildPrivileges(modules []RoleModule) []string {
	privileges := make([]string, 0)
	if len(modules) > 0 {
		for _, module := range modules {
			id := module.ModuleId
			if module.Permissions != 0 {
				id = module.ModuleId + " " + fmt.Sprintf("%X", module.Permissions)
			}
			privileges = append(privileges, id)
		}
	}
	return privileges
}
func getModules(ctx context.Context, db *sql.DB, roleId string, buildParam func(int) string, m map[string]int) ([]RoleModule, error) {
	var modules []RoleModule
	p := buildParam(1)
	query := fmt.Sprintf(`select moduleId, permissions from roleModules where roleId = %s`, p)
	err := q.QueryWithMap(ctx, db, m, &modules, query, roleId)
	return modules, err
}
func toModules(menu string) RoleModule {
	s := strings.Split(menu, " ")
	permission := ActionNone
	if len(s) >= 2 {
		i, err := strconv.ParseInt(s[1], 16, 64)
		if err == nil {
			permission = int32(i)
		}
	}
	p := RoleModule{ModuleId: s[0], Permissions: permission}
	return p
}
func buildPatchRoleStatements(json map[string]interface{}, buildParam func(int) string, modelType reflect.Type) (q.Statements, error) {
	sts := q.NewStatements(true)
	primaryKeyColumns, _ := q.FindPrimaryKeys(modelType)
	jsonColumnMap := q.MakeJsonColumnMap(modelType)
	columnMap := q.JSONToColumns(json, jsonColumnMap)
	sts.Add(q.BuildToPatch("roles", columnMap, primaryKeyColumns, buildParam))

	deleteModules := fmt.Sprintf("delete from rolemodules where roleId = '%s';", json["roleId"])
	sts.Add(deleteModules, nil)

	a, ok := json["privileges"]
	if ok {
		t, _ := a.([]string)
		for i := 0; i < len(t); i++ {
			splitPermission := strings.Fields(t[i])
			insertModules := fmt.Sprintf("insert into rolemodules values ('%s','%s','%s');", buildParam(1), buildParam(2), buildParam(3))
			sts.Add(insertModules, []interface{}{json["roleId"], splitPermission[0], splitPermission[1]})
		}
	}
	return sts, nil
}
