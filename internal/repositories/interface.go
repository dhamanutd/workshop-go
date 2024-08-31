package repositories

import (
	"context"

	postgresql_models "github.com/dhamanutd/golang-clean-architecture/internal/repositories/postgresql/models"
)

type IUserRepository interface {
	Create(ctx context.Context, model *postgresql_models.User) error
	Update(ctx context.Context, model *postgresql_models.User) error
	GetById(ctx context.Context, model *postgresql_models.User, id any) error
	GetAll(ctx context.Context, model *postgresql_models.User, id any) error
}

type IProductRepository interface {
	Create(ctx context.Context, model *postgresql_models.Product) error
	Update(ctx context.Context, model *postgresql_models.Product) error
	GetById(ctx context.Context, model *postgresql_models.Product, id any) error
	GetAll(ctx context.Context, model *postgresql_models.Product, id any) error
}

type IOrderRepository interface {
	Create(ctx context.Context, model *postgresql_models.Order) error
	Update(ctx context.Context, model *postgresql_models.Order) error
	GetById(ctx context.Context, model *postgresql_models.Order, id any) error
	GetAll(ctx context.Context, model *postgresql_models.Order, id any) error
}
