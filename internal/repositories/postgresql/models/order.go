package postgresql_models

import "github.com/google/uuid"

type Order struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID    uuid.UUID `gorm:"column:user_id"`                     // Foreign key to the User table
	User      User      `gorm:"foreignKey:UserID;references:ID"`    // Relation to the User model
	ProductID uuid.UUID `gorm:"column:product_id"`                  // Foreign key to the Product table
	Product   Product   `gorm:"foreignKey:ProductID;references:ID"` // Relation to the Product model
	Amount    float64   `gorm:"column:amount"`
	Status    string    `gorm:"column:status"`
	CreatedAt int64     `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64     `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (o *Order) TableName() string {
	return "orders"
}
