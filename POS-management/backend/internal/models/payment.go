package models

import "time"

// PaymentMethod represents the payment method
type PaymentMethod string

const (
	PaymentMethodQRCode PaymentMethod = "qr_code"
	PaymentMethodCash   PaymentMethod = "cash"
	PaymentMethodCard   PaymentMethod = "card"
)

// PaymentStatus represents the status of a payment
type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusFailed    PaymentStatus = "failed"
)

// Payment represents a payment transaction
type Payment struct {
	ID            uint          `gorm:"primaryKey" json:"id"`
	OrderID       uint          `gorm:"not null" json:"order_id"`
	Amount        float64       `gorm:"type:decimal(10,2);not null" json:"amount"`
	PaymentMethod PaymentMethod `gorm:"type:varchar(20);not null" json:"payment_method"`
	PaymentStatus PaymentStatus `gorm:"type:varchar(20);not null;default:'pending'" json:"payment_status"`
	TransactionID string        `json:"transaction_id"`
	CreatedAt     time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time     `gorm:"autoUpdateTime" json:"updated_at"`

	// Relationships
	Order Order `gorm:"foreignKey:OrderID" json:"order,omitempty"`
}

func (Payment) TableName() string {
	return "payments"
}
