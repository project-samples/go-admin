package fund

import (
	"context"

	"github.com/core-go/core"
)

type FundService interface {
	core.SimpleService
}

func NewFundService(repository core.Repository) FundService {
	return &cinemaService{repository: repository}
}

type cinemaService struct {
	repository core.Repository
}

func (s *cinemaService) Load(context context.Context, id interface{}) (interface{}, error) {
	var cinema Fund
	ok, err := s.repository.LoadAndDecode(context, id, &cinema)
	if !ok {
		return nil, err
	}
	return &cinema, nil
}

func (s *cinemaService) Insert(ctx context.Context, model interface{}) (int64, error) {
	return s.repository.Insert(ctx, model)
}

func (s *cinemaService) Update(ctx context.Context, model interface{}) (int64, error) {
	return s.repository.Update(ctx, model)
}
func (s *cinemaService) Patch(ctx context.Context, model map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, model)
}
func (s *cinemaService) Delete(ctx context.Context, id interface{}) (int64, error) {
	return s.repository.Delete(ctx, id)
}
