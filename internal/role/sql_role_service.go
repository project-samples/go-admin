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
	s "github.com/core-go/sql"
	"github.com/core-go/sql/query"
)

const ActionNone int32 = 0

type SqlRoleService struct {
	db *sql.DB
	sv.GenericService
	search.SearchService
	BuildParam  func(int) string
	CheckDelete string
	modelType   reflect.Type
	Driver      string
}

func NewRoleService(db *sql.DB, checkDelete string) *SqlRoleService {
	var model Role
	tableName := "roles"
	modelType := reflect.TypeOf(model)
	builder := query.NewBuilder(db, tableName, modelType)
	searchService := s.NewSearcherWithQuery(db, modelType, builder.BuildQuery)
	genericService := s.NewWriter(db, tableName, modelType)
	sql := s.ReplaceQueryArgs(s.GetDriver(db), checkDelete)
	buildParam := s.GetBuild(db)
	driver := s.GetDriver(db)
	return &SqlRoleService{db: db, Driver: driver, GenericService: genericService, SearchService: searchService, BuildParam: buildParam, CheckDelete: sql, modelType: modelType}
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
	privileges, er2 := GetPrivileges(ctx, s.db, roleId, s.BuildParam, GetModules)
	if er2 != nil {
		return nil, er2
	}
	role.Privileges = privileges
	return role, nil
}
func (s *SqlRoleService) Insert(ctx context.Context, obj interface{}) (int64, error) {
	sts, err := BuildInsertStatements(ctx, obj, s.Driver, s.BuildParam)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
}
func (s *SqlRoleService) Update(ctx context.Context, obj interface{}) (int64, error) {
	sts, err := BuildUpdateStatements(ctx, obj, s.Driver, s.BuildParam)
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
func BuildInsertStatements(ctx context.Context, obj interface{}, driver string, buildParam func(int) string) (s.Statements, error) {
	role, ok := obj.(*Role)
	if !ok {
		return nil, fmt.Errorf("invalid obj model from request")
	}
	modules, er1 := BuildModules(ctx, role.RoleId, role.Privileges)
	if er1 != nil {
		return nil, er1
	}
	sts := s.NewStatements(true)
	sts.Add(s.BuildToInsert("roles", obj, buildParam))
	query, args, er2 := s.BuildToInsertBatch("roleModules", modules, driver, buildParam)
	if er2 != nil {
		return nil, er2
	}
	sts.Add(query, args)
	return sts, nil
}
func BuildUpdateStatements(ctx context.Context, obj interface{}, driver string, buildParam func(int) string) (s.Statements, error) {
	role, ok := obj.(*Role)
	if !ok {
		return nil, fmt.Errorf("invalid obj model from request")
	}
	modules, err := BuildModules(ctx, role.RoleId, role.Privileges)
	if err != nil {
		return nil, err
	}
	sts := s.NewStatements(true)
	sts.Add(s.BuildToUpdate("roles", obj, buildParam))

	deleteModules := fmt.Sprintf("delete from roleModules where roleId = %s", buildParam(1))
	sts.Add(deleteModules, []interface{}{role.RoleId})
	query, args, er2 := s.BuildToInsertBatch("roleModules", modules, driver, buildParam)
	if er2 != nil {
		return nil, er2
	}
	sts.Add(query, args)
	return sts, nil
}
func BuildDeleteStatements(id interface{}, buildParam func(int) string) (s.Statements, error) {
	roleId, ok := id.(string)
	if !ok {
		return nil, fmt.Errorf("invalid id from request")
	}

	sts := s.NewStatements(false)

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
func GetPrivileges(ctx context.Context, db *sql.DB, roleId string, buildParam func(int) string, getModules func(context.Context, *sql.DB, string, func(int) string) ([]RoleModule, error)) ([]string, error) {
	modules, er1 := getModules(ctx, db, roleId, buildParam)
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
func GetModules(ctx context.Context, db *sql.DB, roleId string, buildParam func(int) string) ([]RoleModule, error) {
	var modules []RoleModule
	p := buildParam(1)
	query := fmt.Sprintf(`select moduleId, permissions from roleModules where roleId = %s`, p)
	err := s.Query(ctx, db, &modules, query, roleId)
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
	sts, err := BuildPatchRoleStatements(ctx, obj, s.BuildParam, s.modelType, s.db)
	if err != nil {
		return 0, err
	}

	return sts.Exec(ctx, s.db)
}

func BuildPatchRoleStatements(ctx context.Context, obj map[string]interface{}, buildParam func(int) string, modelType reflect.Type, db *sql.DB) (s.Statements, error) {
	sts := s.NewStatements(true)
	idcolumNames, idJsonName := s.FindPrimaryKeys(modelType)
	columNames := s.FindJsonName(modelType)
	sts.Add(s.BuildPatch("roles", obj, columNames, idJsonName, idcolumNames, buildParam))

	deleteModules := fmt.Sprintf("delete from rolemodules where roleId = '%s';", obj["roleId"])
	sts.Add(deleteModules, nil)

	if obj["privileges"] != nil {
		a := obj["privileges"]
		t, _ := a.([]string)
		for i := 0; i < len(t); i++ {
			splitPermission := strings.Fields(t[i])
			insertModules := fmt.Sprintf("insert into rolemodules values ('%s','%s','%s');", obj["roleId"], splitPermission[0], splitPermission[1])
			sts.Add(insertModules, nil)
		}
	}

	return sts, nil
}
