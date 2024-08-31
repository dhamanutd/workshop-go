package postgresql_users

import (
	"github.com/dhamanutd/golang-clean-architecture/internal/repositories"
	postgresql_models "github.com/dhamanutd/golang-clean-architecture/internal/repositories/postgresql/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	repositories.BaseRepository[postgresql_models.Product]
}

func New(db *gorm.DB) repositories.IProductRepository {
	return &ProductRepository{
		BaseRepository: repositories.BaseRepository[postgresql_models.Product]{
			DB: db,
		},
	}
}
