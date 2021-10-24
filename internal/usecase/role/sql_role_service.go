package role

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/core-go/search"
	sv "github.com/core-go/service"
	q "github.com/core-go/sql"
	"github.com/core-go/sql/query"
)

const ActionNone int32 = 0

type SqlRoleService struct {
	db *sql.DB
	sv.GenericService
	search.SearchService
	BuildParam       func(int) string
	CheckDelete      string
	modelType        reflect.Type
	Driver           string
	Map              map[string]int
	roleSchema       *q.Schema
	roleModuleSchema *q.Schema
}

func NewRoleService(db *sql.DB, checkDelete string) (*SqlRoleService, error) {
	var model Role
	tableName := "roles"
	modelType := reflect.TypeOf(model)
	builder := query.NewBuilder(db, tableName, modelType)
	searchService, er0 := q.NewSearcherWithQuery(db, modelType, builder.BuildQuery, nil)
	if er0 != nil {
		return nil, er0
	}
	genericService, er1 := q.NewWriter(db, tableName, modelType, nil)
	if er1 != nil {
		return nil, er1
	}
	sql := q.ReplaceQueryArgs(q.GetDriver(db), checkDelete)
	buildParam := q.GetBuild(db)
	driver := q.GetDriver(db)
	var r RoleModule
	subType := reflect.TypeOf(r)
	m, err := q.GetColumnIndexes(subType)
	if err != nil {
		return nil, err
	}
	roleSchema := q.CreateSchema(modelType)
	roleModuleSchema := q.CreateSchema(subType)
	return &SqlRoleService{db: db, Driver: driver, GenericService: genericService, SearchService: searchService, BuildParam: buildParam, CheckDelete: sql, modelType: modelType, Map: m, roleSchema: roleSchema, roleModuleSchema: roleModuleSchema}, nil
}

func (s *SqlRoleService) Load(ctx context.Context, id interface{}) (interface{}, error) {
	roleId, ok := id.(string)
	if !ok {
		return nil, fmt.Errorf("invalid role id")
	}
	rs, er1 := s.GenericService.Load(ctx, roleId)
	if rs == nil && er1 == nil {
		return rs, er1
	}
	role, _ := rs.(*Role)

	privileges, er3 := GetPrivileges(ctx, s.db, roleId, s.BuildParam, GetModules, s.Map)
	if er3 != nil {
		return nil, er3
	}
	role.Privileges = privileges
	return role, nil
}
func (s *SqlRoleService) Insert(ctx context.Context, obj interface{}) (int64, error) {
	sts, err := BuildInsertStatements(ctx, obj, s.Driver, s.BuildParam, s.roleSchema, s.roleModuleSchema)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
}
func (s *SqlRoleService) Update(ctx context.Context, obj interface{}) (int64, error) {
	sts, err := BuildUpdateStatements(ctx, obj, s.Driver, s.BuildParam, s.roleSchema, s.roleModuleSchema)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
}
func (s *SqlRoleService) Delete(ctx context.Context, id interface{}) (int64, error) {
	if len(s.CheckDelete) > 0 {
		exist, er0 := CheckExist(s.db, s.CheckDelete, id)
		if exist || er0 != nil {
			return -1, er0
		}
	}
	sts, er1 := BuildDeleteStatements(id, s.BuildParam)
	if er1 != nil {
		return 0, er1
	}
	return sts.Exec(ctx, s.db)
}
func CheckExist(db *sql.DB, sql string, args ...interface{}) (bool, error) {
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
func BuildInsertStatements(ctx context.Context, obj interface{}, driver string, buildParam func(int) string, roleSchema *q.Schema, roleModuleSchema *q.Schema) (q.Statements, error) {
	role, ok := obj.(*Role)
	if !ok {
		return nil, fmt.Errorf("invalid obj model from request")
	}
	modules, er1 := BuildModules(ctx, role.RoleId, role.Privileges)
	if er1 != nil {
		return nil, er1
	}
	sts := q.NewStatements(true)
	sts.Add(q.BuildToInsert("roles", obj, buildParam, roleSchema))
	query, args, er2 := q.BuildToInsertBatch("roleModules", modules, driver, roleModuleSchema)
	if er2 != nil {
		return nil, er2
	}
	sts.Add(query, args)
	return sts, nil
}
func BuildUpdateStatements(ctx context.Context, obj interface{}, driver string, buildParam func(int) string, roleSchema *q.Schema, roleModuleSchema *q.Schema) (q.Statements, error) {
	role, ok := obj.(*Role)
	if !ok {
		return nil, fmt.Errorf("invalid obj model from request")
	}
	modules, err := BuildModules(ctx, role.RoleId, role.Privileges)
	if err != nil {
		return nil, err
	}
	sts := q.NewStatements(true)
	sts.Add(q.BuildToUpdate("roles", obj, buildParam, roleSchema))

	deleteModules := fmt.Sprintf("delete from roleModules where roleId = %s", buildParam(1))
	sts.Add(deleteModules, []interface{}{role.RoleId})
	query, args, er2 := q.BuildToInsertBatch("roleModules", modules, driver, roleModuleSchema)
	if er2 != nil {
		return nil, er2
	}
	sts.Add(query, args)
	return sts, nil
}
func BuildDeleteStatements(id interface{}, buildParam func(int) string) (q.Statements, error) {
	roleId, ok := id.(string)
	if !ok {
		return nil, fmt.Errorf("invalid id from request")
	}

	sts := q.NewStatements(false)

	deleteModules := fmt.Sprintf("delete from roleModules where roleId = %s", buildParam(1))
	sts.Add(deleteModules, []interface{}{roleId})

	deleteRole := fmt.Sprintf("delete from roles where roleId = %s", buildParam(1))
	sts.Add(deleteRole, []interface{}{roleId})

	return sts, nil
}
func BuildModules(ctx context.Context, roleId string, privileges []string) ([]RoleModule, error) {
	if privileges == nil || len(privileges) <= 0 {
		return nil, nil
	}
	modules := make([]RoleModule, 0)
	for _, p := range privileges {
		m := ToModules(p)
		m.RoleId = roleId
		modules = append(modules, m)
	}
	return modules, nil
}
func GetPrivileges(ctx context.Context, db *sql.DB, roleId string, buildParam func(int) string, getModules func(context.Context, *sql.DB, string, func(int) string, map[string]int) ([]RoleModule, error), m map[string]int) ([]string, error) {
	modules, er1 := getModules(ctx, db, roleId, buildParam, m)
	if er1 != nil {
		return nil, er1
	}
	privileges := BuildPrivileges(modules)
	return privileges, nil
}
func BuildPrivileges(modules []RoleModule) []string {
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
func GetModules(ctx context.Context, db *sql.DB, roleId string, buildParam func(int) string, m map[string]int) ([]RoleModule, error) {
	var modules []RoleModule
	p := buildParam(1)
	query := fmt.Sprintf(`select moduleId, permissions from roleModules where roleId = %s`, p)
	err := q.QueryWithMap(ctx, db, m, &modules, query, roleId)
	return modules, err
}
func ToModules(menu string) RoleModule {
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

func (s *SqlRoleService) Patch(ctx context.Context, obj map[string]interface{}) (int64, error) {
	sts, err := BuildPatchRoleStatements(obj, s.BuildParam, s.modelType)
	if err != nil {
		return 0, err
	}

	return sts.Exec(ctx, s.db)
}

func BuildPatchRoleStatements(json map[string]interface{}, buildParam func(int) string, modelType reflect.Type) (q.Statements, error) {
	sts := q.NewStatements(true)
	primaryKeyColumns, _ := q.FindPrimaryKeys(modelType)
	jsonColumnMap := q.MakeJsonColumnMap(modelType)
	columnMap := q.JSONToColumns(json, jsonColumnMap)
	sts.Add(q.BuildToPatch("roles", columnMap, primaryKeyColumns, buildParam))

	deleteModules := fmt.Sprintf("delete from rolemodules where roleId = '%s';", json["roleId"])
	sts.Add(deleteModules, nil)

	if json["privileges"] != nil {
		a := json["privileges"]
		t, _ := a.([]string)
		for i := 0; i < len(t); i++ {
			splitPermission := strings.Fields(t[i])
			insertModules := fmt.Sprintf("insert into rolemodules values ('%s','%s','%s');", json["roleId"], splitPermission[0], splitPermission[1])
			sts.Add(insertModules, nil)
		}
	}

	return sts, nil
}
