package user

import (
	"database/sql"
	"github.com/core-go/core/unique"
	"net/http"
	"reflect"

	"github.com/core-go/core"
	"github.com/core-go/core/handler/builder"
	v "github.com/core-go/core/validator"
	"github.com/core-go/search/convert"
	q "github.com/core-go/sql"
	"github.com/core-go/sql/query"
	"github.com/core-go/sql/template"
	tb "github.com/core-go/sql/template/builder"
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

func NewUserTransport(db *sql.DB, logError core.Log, templates map[string]*template.Template, tracking builder.TrackingConfig, writeLog core.WriteLog, action *core.ActionConfig) (UserTransport, error) {
	validator, err := v.NewValidator[*User]()
	if err != nil {
		return nil, err
	}
	buildParam := q.GetBuild(db)
	userType := reflect.TypeOf(User{})
	queryUser, err := tb.UseQuery[*UserFilter]("user", templates, &userType, convert.ToMap, buildParam, q.GetSort)
	if err != nil {
		return nil, err
	}
	userSearchBuilder, err := query.NewSearchBuilder[User, *UserFilter](db, queryUser)
	if err != nil {
		return nil, err
	}
	// userValidator := user.NewUserValidator(db, conf.Sql.User, validator.validateFileName)
	userValidator, err := unique.NewUniqueFieldValidator[*User](db, "users", "username", validator.Validate)
	if err != nil {
		return nil, err
	}
	userRepository, er7 := NewUserAdapter(db)
	if er7 != nil {
		return nil, er7
	}
	userService := NewUserService(userRepository)
	userHandler := NewUserHandler(userSearchBuilder.Search, userService, logError, userValidator.Validate, tracking, writeLog, action)
	return userHandler, nil
}
