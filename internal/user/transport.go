package user

import (
	"context"
	"database/sql"
	"net/http"
	"reflect"

	"github.com/core-go/core"
	"github.com/core-go/core/builder"
	"github.com/core-go/core/shortid"
	"github.com/core-go/core/unique"
	v10 "github.com/core-go/core/v10"
	"github.com/core-go/search/convert"
	q "github.com/core-go/sql"
	"github.com/core-go/sql/template"
)

type UserTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetUserByRole(w http.ResponseWriter, r *http.Request)
}

func NewUserTransport(db *sql.DB, logError func(context.Context, string, ...map[string]interface{}), templates map[string]*template.Template, tracking builder.TrackingConfig, writeLog func(context.Context, string, string, bool, string) error, action *core.ActionConfig) (UserTransport, error) {
	validator, err := v10.NewValidator()
	if err != nil {
		return nil, err
	}
	buildParam := q.GetBuild(db)
	userType := reflect.TypeOf(User{})
	queryUser, err := template.UseQuery("user", templates, &userType, convert.ToMap, buildParam, q.GetSort)
	if err != nil {
		return nil, err
	}
	userSearchBuilder, err := q.NewSearchBuilder(db, userType, queryUser)
	if err != nil {
		return nil, err
	}
	// userValidator := user.NewUserValidator(db, conf.Sql.User, validator.validateFileName)
	userValidator := unique.NewUniqueFieldValidator(db, "users", "username", reflect.TypeOf(User{}), validator.Validate)
	userRepository, er7 := NewUserAdapter(db)
	if er7 != nil {
		return nil, er7
	}
	userService := NewUserService(userRepository)
	userHandler := NewUserHandler(userSearchBuilder.Search, userService, shortid.Generate, userValidator.Validate, logError, writeLog, action, tracking)
	return userHandler, nil
}
