package country

import "context"

type CurrencyRepository interface {
	Load(ctx context.Context, id string) (*Country, error)
	Create(ctx context.Context, country *Country) (int64, error)
	Update(ctx context.Context, country *Country) (int64, error)
	Patch(ctx context.Context, country map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}
