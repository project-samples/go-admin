package locale

import (
	"context"
	"github.com/core-go/core"
)

type LocaleService interface {
	Load(ctx context.Context, id string) (*Locale, error)
	Create(ctx context.Context, locale *Locale) (int64, error)
	Update(ctx context.Context, locale *Locale) (int64, error)
	Patch(ctx context.Context, locale map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewLocaleService(repository core.Repository) LocaleService {
	return &localeService{repository: repository}
}

type localeService struct {
	repository core.Repository
}

func (s *localeService) Load(ctx context.Context, id string) (*Locale, error) {
	var locale Locale
	ok, err := s.repository.Get(ctx, id, &locale)
	if !ok {
		return nil, err
	} else {
		return &locale, err
	}
}
func (s *localeService) Create(ctx context.Context, locale *Locale) (int64, error) {
	return s.repository.Insert(ctx, locale)
}
func (s *localeService) Update(ctx context.Context, locale *Locale) (int64, error) {
	return s.repository.Update(ctx, locale)
}
func (s *localeService) Patch(ctx context.Context, locale map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, locale)
}
func (s *localeService) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
