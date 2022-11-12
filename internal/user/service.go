package user

import "context"

type UserService interface {
	Load(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) (int64, error)
	Update(ctx context.Context, user *User) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
	GetUserByRole(ctx context.Context, roleId string) ([]User, error)
}

func NewUserService(repository UserRepository) UserService {
	return &UserUseCase{repository: repository}
}

type UserUseCase struct {
	repository UserRepository
}

func (s *UserUseCase) Load(ctx context.Context, id string) (*User, error) {
	return s.repository.Load(ctx, id)
}
func (s *UserUseCase) Create(ctx context.Context, user *User) (int64, error) {
	return s.repository.Create(ctx, user)
}
func (s *UserUseCase) Update(ctx context.Context, user *User) (int64, error) {
	return s.repository.Update(ctx, user)
}
func (s *UserUseCase) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, user)
}
func (s *UserUseCase) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
func (s *UserUseCase) GetUserByRole(ctx context.Context, roleId string) ([]User, error) {
	return s.repository.GetUserByRole(ctx, roleId)
}