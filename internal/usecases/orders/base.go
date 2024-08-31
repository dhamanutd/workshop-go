package orders

import (
	"github.com/dhamanutd/golang-clean-architecture/internal/repositories"
	"github.com/dhamanutd/golang-clean-architecture/internal/usecases"
)

type OrderUsecase struct {
	repo repositories.IOrderRepository
}

func New(
	repo repositories.IOrderRepository,
) usecases.IOrderUsecase {
	return &OrderUsecase{
		repo: repo,
	}
}
