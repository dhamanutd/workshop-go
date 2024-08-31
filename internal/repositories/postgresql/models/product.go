package postgresql_models

import "github.com/google/uuid"

type Product struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string    `gorm:"column:name"`
	Price     float64   `gorm:"column:price"`
	OrderID   uuid.UUID `gorm:"column:order_id"` // Foreign key to the Order table
	CreatedAt int64     `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64     `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (p *Product) TableName() string {
	return "products"
}
