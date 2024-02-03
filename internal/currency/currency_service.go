package currency

import (
	"context"
	"github.com/core-go/core"
)

type CurrencyService interface {
	Load(ctx context.Context, id string) (*Currency, error)
	Create(ctx context.Context, currency *Currency) (int64, error)
	Update(ctx context.Context, currency *Currency) (int64, error)
	Patch(ctx context.Context, currency map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewCurrencyService(repository core.Repository) CurrencyService {
	return &currencyService{repository: repository}
}

type currencyService struct {
	repository core.Repository
}

func (s *currencyService) Load(ctx context.Context, id string) (*Currency, error) {
	var currency Currency
	ok, err := s.repository.Get(ctx, id, &currency)
	if !ok {
		return nil, err
	} else {
		return &currency, err
	}
}
func (s *currencyService) Create(ctx context.Context, currency *Currency) (int64, error) {
	return s.repository.Insert(ctx, currency)
}
func (s *currencyService) Update(ctx context.Context, currency *Currency) (int64, error) {
	return s.repository.Update(ctx, currency)
}
func (s *currencyService) Patch(ctx context.Context, currency map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, currency)
}
func (s *currencyService) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
