package models

import "time"

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusInProgress OrderStatus = "in_progress"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusPaid       OrderStatus = "paid"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

// Order represents a customer order
type Order struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	TableID     uint        `gorm:"not null" json:"table_id"`
	WaiterID    uint        `gorm:"not null" json:"waiter_id"`
	Status      OrderStatus `gorm:"type:varchar(20);not null;default:'pending'" json:"status"`
	TotalAmount float64     `gorm:"type:decimal(10,2);default:0" json:"total_amount"`
	Notes       string      `json:"notes"`
	CreatedAt   time.Time   `gorm:"autoCreateTime" json:"created_at"`
	CompletedAt *time.Time  `json:"completed_at"`

	// Relationships
	Table      Table       `gorm:"foreignKey:TableID" json:"table,omitempty"`
	Waiter     User        `gorm:"foreignKey:WaiterID" json:"waiter,omitempty"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
	Payments   []Payment   `gorm:"foreignKey:OrderID" json:"payments,omitempty"`
}

func (Order) TableName() string {
	return "orders"
}
