package postgresql_models

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	CreatedAt int64     `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64     `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *User) TableName() string {
	return "users"
}
