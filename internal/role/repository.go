package role

import "context"

type RoleRepository interface {
	Load(ctx context.Context, id string) (*Role, error)
	Create(ctx context.Context, role *Role) (int64, error)
	Update(ctx context.Context, role *Role) (int64, error)
	Patch(ctx context.Context, obj map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
	AssignRole(ctx context.Context, roleId string, users []string) (int64, error)
}
