package country

import (
	"context"
)

type CountryService interface {
	Load(ctx context.Context, id string) (*Country, error)
	Create(ctx context.Context, locale *Country) (int64, error)
	Update(ctx context.Context, locale *Country) (int64, error)
	Patch(ctx context.Context, locale map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}
