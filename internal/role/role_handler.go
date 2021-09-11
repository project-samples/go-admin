package role

import (
	"context"
	"reflect"

	"github.com/core-go/search"
	sv "github.com/core-go/service"
	"github.com/core-go/service/model-builder"
)

type RoleHandler struct {
	*sv.GenericHandler
	*search.SearchHandler
	service RoleService
}

func NewRoleHandler(
		roleService RoleService,
		conf sv.WriterConfig,
		logError func(context.Context, string),
		generateId func(context.Context) (string, error),
		validate func(context.Context, interface{}) ([]sv.ErrorMessage, error),
		tracking builder.TrackingConfig,
		writeLog func(context.Context, string, string, bool, string) error) *RoleHandler {
	searchModelType := reflect.TypeOf(RoleSM{})
	modelType := reflect.TypeOf(Role{})
	searchHandler := search.NewJSONSearchHandler(roleService.Search, modelType, searchModelType, logError, nil)
	modelBuilder := builder.NewDefaultModelBuilderByConfig(generateId, modelType, tracking)
	genericHandler := sv.NewGenericHandlerWithConfig(roleService, modelType, conf.Status, modelBuilder, logError, validate, writeLog)
	return &RoleHandler{GenericHandler: genericHandler, SearchHandler: searchHandler, service: roleService}
}
