package country

import (
	"context"
	"github.com/core-go/core"
)

type CountryService interface {
	Load(ctx context.Context, id string) (*Country, error)
	Create(ctx context.Context, locale *Country) (int64, error)
	Update(ctx context.Context, locale *Country) (int64, error)
	Patch(ctx context.Context, locale map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewCountryService(repository core.Repository) CountryService {
	return &CountryUseCase{repository: repository}
}

type CountryUseCase struct {
	repository core.Repository
}

func (s *CountryUseCase) Load(ctx context.Context, id string) (*Country, error) {
	var locale Country
	ok, err := s.repository.Get(ctx, id, &locale)
	if !ok {
		return nil, err
	} else {
		return &locale, err
	}
}
func (s *CountryUseCase) Create(ctx context.Context, locale *Country) (int64, error) {
	return s.repository.Insert(ctx, locale)
}
func (s *CountryUseCase) Update(ctx context.Context, locale *Country) (int64, error) {
	return s.repository.Update(ctx, locale)
}
func (s *CountryUseCase) Patch(ctx context.Context, locale map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, locale)
}
func (s *CountryUseCase) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
