package user

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	q "github.com/core-go/sql"
)

type userRole struct {
	UserId string `json:"userId,omitempty" gorm:"column:userId;primary_key" bson:"_id,omitempty" validate:"required,max=20,code"`
	RoleId string `json:"roleId,omitempty" gorm:"column:roleId;primary_key" bson:"_id,omitempty" dynamodbav:"roleId,omitempty" firestore:"roleId,omitempty" validate:"max=40"`
}

type UserAdapter struct {
	db            *sql.DB
	driver        string
	BuildParam    func(int) string
	CheckDelete   string
	keys          []string
	jsonColumnMap map[string]string
	Map           map[string]int
	Schema        *q.Schema
	RoleMap       map[string]int
	RoleSchema    *q.Schema
}

func NewUserAdapter(db *sql.DB) (*UserAdapter, error) {
	userMap, userSchema, jsonColumnMap, keys, _, _, buildParam, driver, err := q.Init(reflect.TypeOf(User{}), db)
	if err != nil {
		return nil, err
	}
	roleType := reflect.TypeOf(userRole{})
	userRoleSchema := q.CreateSchema(roleType)
	roleMap, err := q.GetColumnIndexes(roleType)

	return &UserAdapter{
		db:            db,
		driver:        driver,
		BuildParam:    buildParam,
		keys:          keys,
		jsonColumnMap: jsonColumnMap,
		Map:           userMap,
		Schema:        userSchema,
		RoleMap:       roleMap,
		RoleSchema:    userRoleSchema,
	}, err
}

func (s *UserAdapter) Load(ctx context.Context, id string) (*User, error) {
	var users []User
	sql := fmt.Sprintf("select * from users where userId = %s", s.BuildParam(1))
	er1 := q.Query(ctx, s.db, s.Map, &users, sql, id)
	if er1 != nil {
		return nil, er1
	}
	if len(users) == 0 {
		return nil, nil
	}

	var userRoles []userRole
	roles := make([]string, 0)
	query := fmt.Sprintf(`select roleId from userRoles where userId = %s`, s.BuildParam(1))
	err := q.Query(ctx, s.db, s.RoleMap, &userRoles, query, id)
	if err != nil {
		return nil, err
	}
	for _, u := range userRoles {
		roles = append(roles, u.RoleId)
	}
	if len(roles) > 0 {
		users[0].Roles = roles
	}
	return &users[0], nil
}

func (s *UserAdapter) Create(ctx context.Context, user *User) (int64, error) {
	sts, err := buildInsertUserStatements(user, s.driver, s.BuildParam, s.Schema, s.RoleSchema)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
}
func buildInsertUserStatements(user *User, driver string, buildParam func(int) string, userSchema *q.Schema, userRoleSchema *q.Schema) (q.Statements, error) {
	modules, er1 := buildUserModules(user.UserId, user.Roles)
	if er1 != nil {
		return nil, er1
	}
	sts := q.NewStatements(true)
	sts.Add(q.BuildToInsert("users", user, buildParam, userSchema))
	if modules != nil {
		query, args, er2 := q.BuildToInsertBatch("userRoles", modules, driver, userRoleSchema)
		if er2 != nil {
			return sts, er2
		}
		sts.Add(query, args)
	}
	return sts, nil
}

func (s *UserAdapter) Update(ctx context.Context, user *User) (int64, error) {
	sts, err := buildUpdateUserStatements(user, s.driver, s.BuildParam, s.Schema, s.RoleSchema)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
}
func buildUpdateUserStatements(user *User, driver string, buildParam func(int) string, userSchema *q.Schema, userRoleSchema *q.Schema) (q.Statements, error) {
	modules, er1 := buildUserModules(user.UserId, user.Roles)
	if er1 != nil {
		return nil, er1
	}
	sts := q.NewStatements(true)
	sts.Add(q.BuildToUpdate("users", user, buildParam, userSchema))

	deleteModules := fmt.Sprintf("delete from userroles where userId = %s", buildParam(1))
	sts.Add(deleteModules, []interface{}{user.UserId})

	if modules != nil {
		query, args, er2 := q.BuildToInsertBatch("userRoles", modules, driver, userRoleSchema)
		if er2 != nil {
			return sts, er2
		}
		sts.Add(query, args)
	}
	return sts, nil
}

func (s *UserAdapter) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	sts, err := s.buildPatchUserStatements(user)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
}
func (s *UserAdapter) buildPatchUserStatements(json map[string]interface{}) (q.Statements, error) {
	sts := q.NewStatements(true)
	columnMap := q.JSONToColumns(json, s.jsonColumnMap)
	sts.Add(q.BuildToPatch("users", columnMap, s.keys, s.BuildParam))
	if json["roles"] != nil {
		deleteModules := fmt.Sprintf("delete from userRoles where userid = %s", s.BuildParam(1))
		sts.Add(deleteModules, []interface{}{json["userId"]})
		a := json["roles"]
		t, ok := a.([]string)
		if ok {
			for i := 0; i < len(t); i++ {
				insertModules := fmt.Sprintf("insert into userroles values (%s,%s)", s.BuildParam(1), s.BuildParam(2))
				sts.Add(insertModules, []interface{}{json["userId"], t[i]})
			}
		}
	}
	return sts, nil
}

func (s *UserAdapter) Delete(ctx context.Context, id string) (int64, error) {
	if len(s.CheckDelete) > 0 {
		exist, er0 := checkExist(s.db, s.CheckDelete, id)
		if exist || er0 != nil {
			return -1, er0
		}
	}
	sts, er1 := buildDeleteUserStatements(id, s.BuildParam)
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
func buildDeleteUserStatements(id string, buildParam func(int) string) (q.Statements, error) {
	sts := q.NewStatements(false)

	deleteModules := fmt.Sprintf("delete from userroles where userId = %s", buildParam(1))
	sts.Add(deleteModules, []interface{}{id})

	deleteRole := fmt.Sprintf("delete from users where userId = %s", buildParam(1))
	sts.Add(deleteRole, []interface{}{id})

	return sts, nil
}

func buildUserModules(userID string, roles []string) ([]userRole, error) {
	if roles == nil || len(roles) <= 0 {
		return nil, nil
	}
	modules := make([]userRole, 0)
	for _, p := range roles {
		modules = append(modules, userRole{UserId: userID, RoleId: p})
	}
	return modules, nil
}

func (s *UserAdapter) GetUserByRole(ctx context.Context, roleId string) ([]User, error) {
	var users []User
	query := fmt.Sprintf(`select u.* from users u join userroles ur on u.userid = ur.userid where ur.roleid = %s`, s.BuildParam(1))
	err := q.Query(ctx, s.db, s.Map, &users, query, roleId)
	if err != nil {
		return nil, err
	}
	return users, nil
}
