package uam

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/core-go/search"
	sv "github.com/core-go/service"
	s "github.com/core-go/sql"
	"github.com/core-go/sql/query"
	"reflect"
	"strings"
)

type SqlUserService struct {
	db *sql.DB
	sv.GenericService
	search.SearchService
	BuildParam  func(int) string
	CheckDelete string
	modelType   reflect.Type
}

func NewUserService(db *sql.DB) *SqlUserService {
	var model User
	tableName := "users"
	modelType := reflect.TypeOf(model)
	buildParam := s.GetBuild(db)
	builder := query.NewBuilder(db, tableName, modelType, buildParam)
	searchService, genericService := s.NewSearchWriter(db, tableName, modelType, builder.BuildQuery)

	return &SqlUserService{db: db, GenericService: genericService, SearchService: searchService, BuildParam: buildParam, modelType: modelType}
}

func (s *SqlUserService) Load(ctx context.Context, id interface{}) (interface{}, error) {
	userId, ok := id.(string)
	if !ok {
		return nil, fmt.Errorf("invalid user id")
	}
	rs, er1 := s.GenericService.Load(ctx, userId)
	if rs == nil && er1 == nil {
		return rs, er1
	}
	user, _ := rs.(*User)
	roles, er3 := GetRoles(ctx, s.db, userId, s.BuildParam)
	if er3 != nil {
		return nil, er3
	}
	if len(roles) > 0 {
		user.Roles = roles
	}
	return user, nil
}

func GetRoles(ctx context.Context, db *sql.DB, userId string, buildParam func(int) string) ([]string, error) {
	var userRoles []UserRole
	roles := make([]string, 0)
	query := fmt.Sprintf(`select roleId from userRoles where userId = %s`, buildParam(1))
	err := s.Query(ctx, db, &userRoles, query, userId)
	if err != nil {
		return nil, err
	}
	for _, u := range userRoles {
		roles = append(roles, u.RoleId)
	}
	return roles, nil
}

func (s *SqlUserService) Insert(ctx context.Context, obj interface{}) (int64, error) {
	sts, err := BuildInsertUserStatements(ctx, obj, s.BuildParam)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
}

func BuildInsertUserStatements(ctx context.Context, obj interface{}, buildParam func(int) string) (s.Statements, error) {
	user, ok := obj.(*User)
	if !ok {
		return nil, fmt.Errorf("invalid obj model from request")
	}
	modules, err := BuildUserModules(ctx, user.UserId, user.Roles)
	if err != nil {
		return nil, err
	}
	sts := s.NewStatements(true)
	sts.Add(s.BuildToInsert("users", obj, 1, buildParam))
	for i, _ := range modules {
		sts.Add(s.BuildToInsert("userroles", modules[i], 0, buildParam))
	}
	return sts, nil
}

func BuildUserModules(ctx context.Context, userID string, roles []string) ([]UserRole, error) {
	if roles == nil || len(roles) <= 0 {
		return nil, nil
	}
	modules := make([]UserRole, 0)
	for _, p := range roles {
		m := ToUserModules(userID, p)
		m.UserId = userID
		m.RoleId = roles[0]
		modules = append(modules, m)
	}
	return modules, nil
}

func ToUserModules(UserID string, menu string) UserRole {
	s := strings.Split(menu, " ")
	p := UserRole{UserId: UserID, RoleId: s[0]}
	return p
}

func (s *SqlUserService) Update(ctx context.Context, obj interface{}) (int64, error) {
	sts, err := BuildUpdateUserStatements(ctx, obj, s.BuildParam)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
}

func BuildUpdateUserStatements(ctx context.Context, obj interface{}, buildParam func(int) string) (s.Statements, error) {
	user, ok := obj.(*User)
	if !ok {
		return nil, fmt.Errorf("invalid obj model from request")
	}
	modules, err := BuildUserModules(ctx, user.UserId, user.Roles)
	if err != nil {
		return nil, err
	}
	sts := s.NewStatements(true)
	sts.Add(s.BuildToUpdate("users", obj, 0, buildParam))

	deleteModules := fmt.Sprintf("delete from userroles where userId = %s", buildParam(1))
	arg1 := make([]interface{}, 0)
	arg1 = append(arg1, user.UserId)
	sts.Add(deleteModules, arg1)

	for i, _ := range modules {
		sts.Add(s.BuildToInsert("userroles", modules[i], 0, buildParam))
	}
	return sts, nil
}

func (s *SqlUserService) Delete(ctx context.Context, id interface{}) (int64, error) {
	if len(s.CheckDelete) > 0 {
		exist, er0 := CheckExist(s.db, s.CheckDelete, id)
		if exist || er0 != nil {
			return -1, er0
		}
	}
	sts, er1 := BuildDeleteUserStatements(id, s.BuildParam)
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
func BuildDeleteUserStatements(id interface{}, buildParam func(int) string) (s.Statements, error) {
	roleId, ok := id.(string)
	if !ok {
		return nil, fmt.Errorf("invalid id from request")
	}

	sts := s.NewStatements(false)

	deleteModules := fmt.Sprintf("delete from userroles where userId = %s", buildParam(1))
	arg1 := make([]interface{}, 0)
	arg1 = append(arg1, roleId)
	sts.Add(deleteModules, arg1)

	deleteRole := fmt.Sprintf("delete from users where userId = %s", buildParam(1))
	arg2 := make([]interface{}, 0)
	arg2 = append(arg2, roleId)
	sts.Add(deleteRole, arg2)

	return sts, nil
}

/*
func (s *SqlUserService) Patch(ctx context.Context, obj map[string]interface{}) (int64, error) {
	sts, err := BuildPatchUserStatements(ctx, obj, s.BuildParam, s.modelType)
	if err != nil {
		return 0, err
	}

	return sts.Exec(ctx, s.db)
}
*/
func BuildPatchUserStatements(ctx context.Context, obj map[string]interface{}, buildParam func(int) string, modelType reflect.Type) (s.Statements, error) {
	sts := s.NewStatements(true)
	idcolumNames, idJsonName := s.FindPrimaryKeys(modelType)
	columNames := s.FindJsonName(modelType)
	sts.Add(s.BuildPatch("users", obj, columNames, idJsonName, idcolumNames, buildParam))
	if obj["roles"] != nil {
		deleteModules := fmt.Sprintf("delete from userRoles where userid = '%s';", obj["userId"])
		sts.Add(deleteModules, nil)
		a := obj["roles"]
		t, _ := a.([]string)
		for i := 0; i < len(t); i++ {
			insertModules := fmt.Sprintf("insert into userroles values ('%s','%s');", obj["userId"], t[i])
			sts.Add(insertModules, nil)
		}
	}
	return sts, nil
}
