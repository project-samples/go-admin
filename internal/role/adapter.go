package role

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	q "github.com/core-go/sql"
	"reflect"
	"strconv"
	"strings"
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
	query1 := fmt.Sprintf("select * from roles where roleId = %s", s.BuildParam(1))
	er1 := q.Query(ctx, s.db, s.Map, &roles, query1, roleId)
	if er1 != nil {
		return nil, er1
	}
	if len(roles) == 0 {
		return nil, nil
	}
	role := roles[0]
	var modules []roleModule
	query2 := fmt.Sprintf(`select moduleId, permissions from roleModules where roleId = %s`, s.BuildParam(1))
	er3 := q.Query(ctx, s.db, s.ModuleMap, &modules, query2, roleId)
	if er3 != nil {
		return nil, er3
	}
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
	role.Privileges = privileges
	return &role, nil
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
func (s *RoleAdapter) Create(ctx context.Context, role *Role) (int64, error) {
	modules, er1 := buildModules(role.RoleId, role.Privileges)
	if er1 != nil {
		return -1, er1
	}
	sts := q.NewStatements(true)
	sts.Add(q.BuildToInsert("roles", role, s.BuildParam, s.Schema))
	if modules != nil {
		query, args, er2 := q.BuildToInsertBatch("roleModules", modules, s.Driver, s.ModuleSchema)
		if er2 != nil {
			return -1, er2
		}
		sts.Add(query, args)
	}

	return sts.Exec(ctx, s.db)
}

func (s *RoleAdapter) Update(ctx context.Context, role *Role) (int64, error) {
	modules, err := buildModules(role.RoleId, role.Privileges)
	if err != nil {
		return -1, err
	}
	sts := q.NewStatements(true)
	sts.Add(q.BuildToUpdate("roles", role, s.BuildParam, s.Schema))

	deleteModules := fmt.Sprintf("delete from roleModules where roleId = %s", s.BuildParam(1))
	sts.Add(deleteModules, []interface{}{role.RoleId})
	if modules != nil {
		query, args, er2 := q.BuildToInsertBatch("roleModules", modules, s.Driver, s.ModuleSchema)
		if er2 != nil {
			return -1, er2
		}
		sts.Add(query, args)
	}

	return sts.Exec(ctx, s.db)
}

func (s *RoleAdapter) Patch(ctx context.Context, role map[string]interface{}) (int64, error) {
	objId, ok := role["roleId"]
	if !ok {
		return -1, errors.New("roleId must be in payload")
	}
	roleId, ok2 := objId.(string)
	if !ok2 {
		return -1, errors.New("roleId must be a string")
	}
	var privileges []string
	var ok4 bool
	objPrivileges, ok3 := role["privileges"]
	if ok3 {
		privileges, ok4 = objPrivileges.([]string)
	}
	var sts q.Statements
	if ok4 && len(role) > 2 || !ok4 && len(role) > 1 {
		sts = q.NewStatements(true)
		columnMap := q.JSONToColumns(role, s.jsonColumnMap)
		sts.Add(q.BuildToPatch("roles", columnMap, s.keys, s.BuildParam))
	} else {
		sts = q.NewStatements(false)
	}

	deleteModules := fmt.Sprintf("delete from rolemodules where roleId = %s", s.BuildParam(1))
	sts.Add(deleteModules, []interface{}{roleId})

	if ok4 {
		modules, err := buildModules(roleId, privileges)
		if err != nil {
			return -1, err
		}
		query, args, er2 := q.BuildToInsertBatch("roleModules", modules, s.Driver, s.ModuleSchema)
		if er2 != nil {
			return -1, er2
		}
		sts.Add(query, args)
	}
	return sts.Exec(ctx, s.db)
}

func (s *RoleAdapter) Delete(ctx context.Context, id string) (int64, error) {
	if len(s.CheckDelete) > 0 {
		exist, er0 := checkExist(s.db, s.CheckDelete, id)
		if exist || er0 != nil {
			return -1, er0
		}
	}
	sts := q.NewStatements(false)

	deleteModules := fmt.Sprintf("delete from roleModules where roleId = %s", s.BuildParam(1))
	sts.Add(deleteModules, []interface{}{id})

	deleteRole := fmt.Sprintf("delete from roles where roleId = %s", s.BuildParam(1))
	sts.Add(deleteRole, []interface{}{id})

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

func (s *RoleAdapter) AssignRole(ctx context.Context, roleId string, users []string) (int64, error) {
	modules, err := buildRoleUser(roleId, users)
	if err != nil {
		return -1, err
	}
	sts := q.NewStatements(true)

	deleteModules := fmt.Sprintf("delete from userroles where roleId = %s", s.BuildParam(1))
	sts.Add(deleteModules, []interface{}{roleId})
	if modules != nil {
		query, args, er2 := q.BuildToInsertBatch("userRoles", modules, s.Driver, s.UserSchema)
		if er2 != nil {
			return -1, er2
		}
		sts.Add(query, args)
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
