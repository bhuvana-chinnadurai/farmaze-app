package model

import (
	"time"

	"github.com/google/uuid"
)

type B2BClient struct {
	ID          uuid.UUID `json:"id"`
	CompanyName string    `json:"company_name"`
	ContactName string    `json:"contact_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
}

type OrderProduct struct {
	OrderID   uuid.UUID `json:"order_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
}

type Order struct {
	ID         uuid.UUID      `json:"id"`
	ClientID   uuid.UUID      `json:"client_id"`
	Products   []OrderProduct `json:"products"`
	TotalPrice float64        `json:"total_price"`
	CreatedAt  time.Time      `json:"created_at"`
	Status     uuid.UUID      `json:"status"`
}

type Product struct {
	ID                uuid.UUID `json:"id"`
	Name              string    `json:"name"`
	Price             float64   `json:"price"`
	Description       string    `json:"description"`
	AvailableQuantity int       `json:"available_quantity"`
	Category          string    `json:"category"`
}

// NotFoundError represents a custom error indicating a resource was not found.
type NotFoundError struct {
	Message string // Additional message to describe the error.
}

// Error returns the error message.
func (e NotFoundError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return "Not Found"
}
