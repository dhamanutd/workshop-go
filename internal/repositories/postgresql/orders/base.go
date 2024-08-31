package postgresql_users

import (
	"github.com/dhamanutd/golang-clean-architecture/internal/repositories"
	postgresql_models "github.com/dhamanutd/golang-clean-architecture/internal/repositories/postgresql/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	repositories.BaseRepository[postgresql_models.Order]
}

func New(db *gorm.DB) repositories.IOrderRepository {
	return &OrderRepository{
		BaseRepository: repositories.BaseRepository[postgresql_models.Order]{
			DB: db,
		},
	}
}
