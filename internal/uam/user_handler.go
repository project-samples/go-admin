package uam

import (
	"context"
	"reflect"

	"github.com/core-go/search"
	sv "github.com/core-go/service"
	"github.com/core-go/service/model-builder"
)

type UserHandler struct {
	*sv.GenericHandler
	*search.SearchHandler
	service UserService
}

func NewUserHandler(
		userService UserService,
		conf sv.WriterConfig,
		logError func(context.Context, string),
		generateId func(context.Context) (string, error),
		validate func(context.Context, interface{}) ([]sv.ErrorMessage, error),
		tracking builder.TrackingConfig,
		writeLog func(context.Context, string, string, bool, string) error) *UserHandler {
	searchModelType := reflect.TypeOf(UserSM{})
	modelType := reflect.TypeOf(User{})
	searchHandler := search.NewJSONSearchHandler(userService.Search, modelType, searchModelType, logError, nil)
	modelBuilder := builder.NewDefaultModelBuilderByConfig(generateId, modelType, tracking)
	genericHandler := sv.NewGenericHandlerWithConfig(userService, modelType, conf.Status, modelBuilder, logError, validate, writeLog)
	return &UserHandler{genericHandler, searchHandler, userService}
}
