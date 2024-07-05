package locale

import (
	"context"
	"database/sql"
	"github.com/core-go/core/tx"
)

type LocaleService interface {
	Load(ctx context.Context, id string) (*Locale, error)
	Create(ctx context.Context, locale *Locale) (int64, error)
	Update(ctx context.Context, locale *Locale) (int64, error)
	Patch(ctx context.Context, locale map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewLocaleService(repository LocaleRepository) LocaleService {
	return &LocaleUseCase{repository: repository}
}

type LocaleUseCase struct {
	db         *sql.DB
	repository LocaleRepository
}

func (s *LocaleUseCase) Load(ctx context.Context, id string) (*Locale, error) {
	return s.repository.Load(ctx, id)
}
func (s *LocaleUseCase) Create(ctx context.Context, locale *Locale) (int64, error) {
	return tx.Execute(ctx, s.db, func(ctx context.Context) (int64, error) {
		return s.repository.Create(ctx, locale)
	})
}
func (s *LocaleUseCase) Update(ctx context.Context, locale *Locale) (int64, error) {
	return tx.Execute(ctx, s.db, func(ctx context.Context) (int64, error) {
		return s.repository.Update(ctx, locale)
	})
}
func (s *LocaleUseCase) Patch(ctx context.Context, locale map[string]interface{}) (int64, error) {
	return tx.Execute(ctx, s.db, func(ctx context.Context) (int64, error) {
		return s.repository.Patch(ctx, locale)
	})
}
func (s *LocaleUseCase) Delete(ctx context.Context, id string) (int64, error) {
	return tx.Execute(ctx, s.db, func(ctx context.Context) (int64, error) {
		return s.repository.Delete(ctx, id)
	})
}
