package locale

import "context"

type LocaleRepository interface {
	Load(ctx context.Context, id string) (*Locale, error)
	Create(ctx context.Context, locale *Locale) (int64, error)
	Update(ctx context.Context, locale *Locale) (int64, error)
	Patch(ctx context.Context, locale map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}
