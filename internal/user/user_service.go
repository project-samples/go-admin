package user

import (
	"github.com/core-go/search"
	sv "github.com/core-go/service"
)

type UserService interface {
	sv.GenericService
	search.SearchService
}
