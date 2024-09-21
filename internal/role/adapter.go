package role

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	q "github.com/core-go/sql"
)

const ActionNone int32 = 0

type userRole struct {
	UserId string `json:"userId,omitempty" gorm:"column:userId;primary_key" bson:"_id,omitempty" validate:"required,max=20,code"`
	RoleId string `json:"roleId,omitempty" gorm:"column:roleId;primary_key" bson:"_id,omitempty" dynamodbav:"roleId,omitempty" firestore:"roleId,omitempty" validate:"max=40"`
}

type roleModule struct {
	RoleId      string `json:"roleId,omitempty" gorm:"column:roleId" bson:"roleId,omitempty" dynamodbav:"roleId,omitempty" firestore:"roleId,omitempty" validate:"required"`
	ModuleId    string `json:"moduleId,omitempty" gorm:"column:moduleId" bson:"moduleId,omitempty" dynamodbav:"moduleId,omitempty" firestore:"moduleId,omitempty" validate:"required"`
	Permissions int32  `json:"permissions,omitempty" gorm:"column:permissions" bson:"permissions,omitempty" dynamodbav:"permissions,omitempty" firestore:"permissions,omitempty" validate:"required"`
}

type RoleAdapter struct {
	db            *sql.DB
	Driver        string
	BuildParam    func(int) string
	CheckDelete   string
	keys          []string
	jsonColumnMap map[string]string
	Map           map[string]int
	Schema        *q.Schema
	ModuleMap     map[string]int
	ModuleSchema  *q.Schema
	UserSchema    *q.Schema
}

func NewRoleAdapter(db *sql.DB, checkDelete string) (*RoleAdapter, error) {
	sql := q.ReplaceQueryArgs(q.GetDriver(db), checkDelete)
	roleMap, roleSchema, jsonColumnMap, keys, _, _, buildParam, driver, err := q.Init(reflect.TypeOf(Role{}), db)
	if err != nil {
		return nil, err
	}
	userRoleSchema := q.CreateSchema(reflect.TypeOf(userRole{}))

	moduleType := reflect.TypeOf(roleModule{})
	roleModuleSchema := q.CreateSchema(moduleType)
	moduleMap, err := q.GetColumnIndexes(moduleType)

	return &RoleAdapter{
			db:            db,
			Driver:        driver,
			BuildParam:    buildParam,
			CheckDelete:   sql,
			Map:           roleMap,
			Schema:        roleSchema,
			jsonColumnMap: jsonColumnMap,
			keys:          keys,
			ModuleMap:     moduleMap,
			ModuleSchema:  roleModuleSchema,
			UserSchema:    userRoleSchema,
		},
		err
}

func (s *RoleAdapter) Load(ctx context.Context, roleId string) (*Role, error) {
	var roles []Role
	sql := fmt.Sprintf("select * from roles where roleId = %s", s.BuildParam(1))
	er1 := q.Query(ctx, s.db, s.Map, &roles, sql, roleId)
	if er1 != nil {
		return nil, er1
	}
	if len(roles) == 0 {
		return nil, nil
	}
	role := roles[0]
	privileges, er3 := getPrivileges(ctx, s.db, roleId, s.BuildParam, getModules, s.ModuleMap)
	if er3 != nil {
		return nil, er3
	}
	role.Privileges = privileges
	return &role, nil
}
func getPrivileges(ctx context.Context, db *sql.DB, roleId string, buildParam func(int) string, getModules func(context.Context, *sql.DB, string, func(int) string, map[string]int) ([]roleModule, error), m map[string]int) ([]string, error) {
	modules, er1 := getModules(ctx, db, roleId, buildParam, m)
	if er1 != nil {
		return nil, er1
	}
	privileges := buildPrivileges(modules)
	return privileges, nil
}
func buildPrivileges(modules []roleModule) []string {
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
func getModules(ctx context.Context, db *sql.DB, roleId string, buildParam func(int) string, m map[string]int) ([]roleModule, error) {
	var modules []roleModule
	p := buildParam(1)
	query := fmt.Sprintf(`select moduleId, permissions from roleModules where roleId = %s`, p)
	err := q.Query(ctx, db, m, &modules, query, roleId)
	return modules, err
}

func (s *RoleAdapter) Create(ctx context.Context, role *Role) (int64, error) {
	sts, err := buildInsertStatements(role, s.Driver, s.BuildParam, s.Schema, s.ModuleSchema)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
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

func (s *RoleAdapter) Update(ctx context.Context, role *Role) (int64, error) {
	sts, err := buildUpdateStatements(role, s.Driver, s.BuildParam, s.Schema, s.ModuleSchema)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
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

func (s *RoleAdapter) Patch(ctx context.Context, obj map[string]interface{}) (int64, error) {
	roleId, ok := obj["roleId"]
	if !ok {
		return -1, errors.New("roleId must be in payload")
	}
	sroleId, ok2 := roleId.(string)
	if !ok2 {
		return -1, errors.New("roleId must be a string")
	}
	sts, err := s.buildPatchRoleStatements(sroleId, obj)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
}
func (s *RoleAdapter) buildPatchRoleStatements(roleId string, json map[string]interface{}) (q.Statements, error) {
	sts := q.NewStatements(true)

	columnMap := q.JSONToColumns(json, s.jsonColumnMap)
	sts.Add(q.BuildToPatch("roles", columnMap, s.keys, s.BuildParam))

	deleteModules := fmt.Sprintf("delete from rolemodules where roleId = %s", s.BuildParam(1))
	sts.Add(deleteModules, []interface{}{roleId})

	a, ok := json["privileges"]
	if ok {
		t, _ := a.([]string)
		modules, err := buildModules(roleId, t)
		if err != nil {
			return nil, err
		}
		query, args, er2 := q.BuildToInsertBatch("roleModules", modules, s.Driver, s.ModuleSchema)
		if er2 != nil {
			return nil, er2
		}
		sts.Add(query, args)
	}
	return sts, nil
}

func (s *RoleAdapter) Delete(ctx context.Context, id string) (int64, error) {
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
func buildDeleteStatements(roleId string, buildParam func(int) string) (q.Statements, error) {
	sts := q.NewStatements(false)

	deleteModules := fmt.Sprintf("delete from roleModules where roleId = %s", buildParam(1))
	sts.Add(deleteModules, []interface{}{roleId})

	deleteRole := fmt.Sprintf("delete from roles where roleId = %s", buildParam(1))
	sts.Add(deleteRole, []interface{}{roleId})

	return sts, nil
}

func (s *RoleAdapter) AssignRole(ctx context.Context, roleId string, users []string) (int64, error) {
	sts, err := buildAssignRoleStatements(roleId, users, s.Driver, s.BuildParam, s.UserSchema)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
}
func buildRoleUser(roleId string, users []string) ([]userRole, error) {
	if users == nil || len(users) <= 0 {
		return nil, nil
	}
	modules := make([]userRole, 0)
	for _, u := range users {
		modules = append(modules, userRole{UserId: u, RoleId: roleId})
	}
	return modules, nil
}
func buildAssignRoleStatements(roleId string, users []string, driver string, buildParam func(int) string, userRoleSchema *q.Schema) (q.Statements, error) {
	modules, err := buildRoleUser(roleId, users)
	if err != nil {
		return nil, err
	}
	sts := q.NewStatements(true)

	deleteModules := fmt.Sprintf("delete from userroles where roleId = %s", buildParam(1))
	sts.Add(deleteModules, []interface{}{roleId})
	if modules != nil {
		query, args, er2 := q.BuildToInsertBatch("userRoles", modules, driver, userRoleSchema)
		if er2 != nil {
			return nil, er2
		}
		sts.Add(query, args)
	}
	return sts, nil
}

func buildModules(roleId string, privileges []string) ([]roleModule, error) {
	if privileges == nil || len(privileges) <= 0 {
		return nil, nil
	}
	modules := make([]roleModule, 0)
	for _, p := range privileges {
		m := toModules(p)
		m.RoleId = roleId
		modules = append(modules, m)
	}
	return modules, nil
}
func toModules(menu string) roleModule {
	s := strings.Split(menu, " ")
	permission := ActionNone
	if len(s) >= 2 {
		i, err := strconv.ParseInt(s[1], 16, 64)
		if err == nil {
			permission = int32(i)
		}
	}
	p := roleModule{ModuleId: s[0], Permissions: permission}
	return p
}
