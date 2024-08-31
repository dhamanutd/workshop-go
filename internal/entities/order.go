package entities

import (
	"fmt"
	"time"

	postgresql_models "github.com/dhamanutd/golang-clean-architecture/internal/repositories/postgresql/models"
	"github.com/google/uuid"
)

// CreateOrderRequest represents the data required to create an order
type CreateOrderRequest struct {
	UserID   uuid.UUID          `json:"user_id" validate:"required,uuid"`                                   // ID of the user placing the order
	Products []OrderProductItem `json:"products" validate:"required"`                                       // List of products to be included in the order
	Amount   float64            `json:"amount" validate:"required,gt=0"`                                    // Total amount for the order
	Status   string             `json:"status" validate:"required,oneof='pending' 'completed' 'cancelled'"` // Status of the order
}

// OrderProductItem represents a single product in the order
type OrderProductItem struct {
	ProductID string  `json:"product_id" validate:"required"`    // ID of the product
	Quantity  int     `json:"quantity" validate:"required,gt=0"` // Quantity of the product in the order
	Price     float64 `json:"price" validate:"required,gt=0"`    // Price of the product at the time of order
}

// ConvertCreateOrderRequestToOrder maps CreateOrderRequest to the Order model.
func (r *CreateOrderRequest) ToModel() postgresql_models.Order {
	orderId := uuid.New()
	products := make([]postgresql_models.Product, len(r.Products))
	for i, p := range r.Products {
		products[i] = postgresql_models.Product{
			ID:        uuid.New(), // Assume a function to generate a unique product ID
			Name:      "",         // Name might not be part of the request, fill it later if needed
			Price:     p.Price,
			OrderID:   orderId, // This will be set when the order is created
			CreatedAt: time.Now().UnixMilli(),
			UpdatedAt: time.Now().UnixMilli(),
		}
	}

	fmt.Printf("akjnwdakwjdnawknd %v", r)

	return postgresql_models.Order{
		ID:        orderId, // Assume a function to generate a unique order ID
		UserID:    r.UserID,
		Amount:    r.Amount,
		Status:    r.Status,
		Product:   products[0],
		CreatedAt: time.Now().UnixMilli(),
		UpdatedAt: time.Now().UnixMilli(),
	}
}
