package postgresql_users

import (
	"github.com/dhamanutd/golang-clean-architecture/internal/repositories"
	postgresql_models "github.com/dhamanutd/golang-clean-architecture/internal/repositories/postgresql/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	repositories.BaseRepository[postgresql_models.User]
}

func New(db *gorm.DB) repositories.IUserRepository {
	return &UserRepository{
		BaseRepository: repositories.BaseRepository[postgresql_models.User]{
			DB: db,
		},
	}
}
