package orders

import (
	"context"

	"github.com/dhamanutd/golang-clean-architecture/internal/entities"
)

func (s *OrderUsecase) Create(ctx context.Context, req *entities.CreateOrderRequest) error {
	model := req.ToModel()
	err := s.repo.Create(ctx, &model)
	if err != nil {
		return err
	}

	return nil
}
