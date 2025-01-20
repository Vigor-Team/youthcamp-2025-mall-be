package models

import (
	"time"

	"github.com/google/uuid" // 导入 uuid 包以生成唯一 ID
)

// Payment represents a payment record.
type Payment struct {
	ID             string    `json:"id" db:"id"`
	OrderID        *string   `json:"order_id" db:"order_id"`
	PaymentMethod  string    `json:"payment_method" db:"payment_method"`
	TotalAmount    float64   `json:"total_amount" db:"total_amount"`
	Description    *string   `json:"description" db:"description"`
	Status         string    `json:"status" db:"status"`
	PaymentGateway string    `json:"payment_gateway" db:"payment_gateway"`
	PaymentID      *string   `json:"payment_id" db:"payment_id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

func NewPayment() *Payment {
	return &Payment{
		ID: uuid.New().String(),
	}
}
