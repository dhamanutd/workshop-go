package orders_controller

import (
	"github.com/dhamanutd/golang-clean-architecture/api/rest/controller"
	"github.com/dhamanutd/golang-clean-architecture/internal/usecases"
	"github.com/dhamanutd/golang-clean-architecture/pkg/validator"
)

type OrderController struct {
	usecase   usecases.IOrderUsecase
	validator *validator.Validate
}

func New(usecase usecases.IOrderUsecase) controller.V1OrderController {
	return &OrderController{
		usecase: usecase,
	}
}
