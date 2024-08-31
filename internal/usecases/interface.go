package usecases

import (
	"context"

	"github.com/dhamanutd/golang-clean-architecture/internal/entities"
)

type IOrderUsecase interface {
	Create(ctx context.Context, req *entities.CreateOrderRequest) error
}
