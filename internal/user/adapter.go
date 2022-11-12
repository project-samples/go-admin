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
	db             *sql.DB
	driver         string
	BuildParam     func(int) string
	CheckDelete    string
	MapUser        map[string]int
	Map            map[string]int
	modelType      reflect.Type
	userSchema     *q.Schema
	userRoleSchema *q.Schema
}

func NewUserRepository(db *sql.DB) (*UserAdapter, error) {
	modelType := reflect.TypeOf(User{})
	buildParam := q.GetBuild(db)
	mu, err := q.GetColumnIndexes(modelType)
	if err != nil {
		return nil, err
	}
	var r userRole
	subType := reflect.TypeOf(r)
	m, err := q.GetColumnIndexes(subType)
	if err != nil {
		return nil, err
	}
	userSchema := q.CreateSchema(modelType)
	userRoleSchema := q.CreateSchema(subType)
	driver := q.GetDriver(db)
	return &UserAdapter{
		db:             db,
		driver:         driver,
		BuildParam:     buildParam,
		modelType:      modelType,
		MapUser:        mu,
		Map:            m,
		userSchema:     userSchema,
		userRoleSchema: userRoleSchema,
	}, nil
}

func (s *UserAdapter) Load(ctx context.Context, id string) (*User, error) {
	var users []User
	sql := fmt.Sprintf("select * from users where userId = %s", s.BuildParam(1))
	er1 := q.Query(ctx, s.db, s.MapUser, &users, sql, id)
	if er1 != nil {
		return nil, er1
	}
	if len(users) == 0 {
		return nil, nil
	}
	user := users[0]
	roles, er2 := getRoles(ctx, s.db, id, s.BuildParam, s.Map)
	if er2 != nil {
		return nil, er2
	}
	if len(roles) > 0 {
		user.Roles = roles
	}
	return &user, nil
}

func (s *UserAdapter) Create(ctx context.Context, user *User) (int64, error) {
	sts, err := buildInsertUserStatements(user, s.driver, s.BuildParam, s.userSchema, s.userRoleSchema)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
}

func (s *UserAdapter) Update(ctx context.Context, user *User) (int64, error) {
	sts, err := buildUpdateUserStatements(user, s.driver, s.BuildParam, s.userSchema, s.userRoleSchema)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
}

func (s *UserAdapter) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	sts, err := buildPatchUserStatements(user, s.BuildParam, s.modelType)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
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

func getRoles(ctx context.Context, db *sql.DB, userId string, buildParam func(int) string, m map[string]int) ([]string, error) {
	var userRoles []userRole
	roles := make([]string, 0)
	query := fmt.Sprintf(`select roleId from userRoles where userId = %s`, buildParam(1))
	err := q.Query(ctx, db, m, &userRoles, query, userId)
	if err != nil {
		return nil, err
	}
	for _, u := range userRoles {
		roles = append(roles, u.RoleId)
	}
	return roles, nil
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

func buildPatchUserStatements(json map[string]interface{}, buildParam func(int) string, modelType reflect.Type) (q.Statements, error) {
	sts := q.NewStatements(true)
	primaryKeyColumns, _ := q.FindPrimaryKeys(modelType)
	jsonColumnMap := q.MakeJsonColumnMap(modelType)
	columnMap := q.JSONToColumns(json, jsonColumnMap)
	sts.Add(q.BuildToPatch("users", columnMap, primaryKeyColumns, buildParam))
	if json["roles"] != nil {
		deleteModules := fmt.Sprintf("delete from userRoles where userid = '%s';", buildParam(1))
		sts.Add(deleteModules, []interface{}{json["userId"]})
		a := json["roles"]
		t, ok := a.([]string)
		if ok {
			for i := 0; i < len(t); i++ {
				insertModules := fmt.Sprintf("insert into userroles values ('%s','%s');", buildParam(1), buildParam(2))
				sts.Add(insertModules, []interface{}{json["userId"], t[i]})
			}
		}
	}
	return sts, nil
}

func (s *UserAdapter) GetUserByRole(ctx context.Context, roleId string) ([]User, error) {
	return getUserByRole(ctx, s.db, roleId, s.BuildParam, s.MapUser)
}

func getUserByRole(ctx context.Context, db *sql.DB, roleId string, buildParam func(int) string, m map[string]int) ([]User, error) {
	var users []User
	query := fmt.Sprintf(`select u.* from users u join userroles ur on u.userid = ur.userid where ur.roleid = %s`, buildParam(1))
	err := q.Query(ctx, db, m, &users, query, roleId)
	if err != nil {
		return nil, err
	}
	return users, nil
}
