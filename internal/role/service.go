package role

import "context"

type RoleService interface {
	Load(ctx context.Context, id string) (*Role, error)
	Create(ctx context.Context, role *Role) (int64, error)
	Update(ctx context.Context, role *Role) (int64, error)
	Patch(ctx context.Context, role map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
	AssignRole(ctx context.Context, roleId string, users []string) (int64, error)
}

func NewRoleService(repository RoleRepository) RoleService {
	return &RoleUseCase{repository: repository}
}

type RoleUseCase struct {
	repository RoleRepository
}

func (s *RoleUseCase) Load(ctx context.Context, id string) (*Role, error) {
	return s.repository.Load(ctx, id)
}
func (s *RoleUseCase) Create(ctx context.Context, role *Role) (int64, error) {
	return s.repository.Create(ctx, role)
}
func (s *RoleUseCase) Update(ctx context.Context, role *Role) (int64, error) {
	return s.repository.Update(ctx, role)
}
func (s *RoleUseCase) Patch(ctx context.Context, role map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, role)
}
func (s *RoleUseCase) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
func (s *RoleUseCase) AssignRole(ctx context.Context, roleId string, users []string) (int64, error) {
	return s.repository.AssignRole(ctx, roleId, users)
}
