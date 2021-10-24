package role

import (
	"github.com/core-go/search"
	sv "github.com/core-go/service"
)

type RoleService interface {
	sv.GenericService
	search.SearchService
}
