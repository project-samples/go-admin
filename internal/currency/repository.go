package currency

import "context"

type CurrencyRepository interface {
	Load(ctx context.Context, id string) (*Currency, error)
	Create(ctx context.Context, currency *Currency) (int64, error)
	Update(ctx context.Context, currency *Currency) (int64, error)
	Patch(ctx context.Context, currency map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}
