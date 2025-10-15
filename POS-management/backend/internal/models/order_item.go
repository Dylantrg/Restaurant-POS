package models

import "time"

// OrderItemStatus represents the status of an order item
type OrderItemStatus string

const (
	OrderItemStatusPending    OrderItemStatus = "pending"
	OrderItemStatusInProgress OrderItemStatus = "in_progress"
	OrderItemStatusCompleted  OrderItemStatus = "completed"
)

// OrderItem represents an item in an order
type OrderItem struct {
	ID                  uint            `gorm:"primaryKey" json:"id"`
	OrderID             uint            `gorm:"not null" json:"order_id"`
	MenuItemID          uint            `gorm:"not null" json:"menu_item_id"`
	Quantity            int             `gorm:"not null" json:"quantity"`
	UnitPrice           float64         `gorm:"type:decimal(10,2);not null" json:"unit_price"`
	TotalPrice          float64         `gorm:"type:decimal(10,2);not null" json:"total_price"`
	SpecialInstructions string          `json:"special_instructions"`
	Status              OrderItemStatus `gorm:"type:varchar(20);not null;default:'pending'" json:"status"`
	CreatedAt           time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time       `gorm:"autoUpdateTime" json:"updated_at"`

	// Relationships
	Order    Order    `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	MenuItem MenuItem `gorm:"foreignKey:MenuItemID" json:"menu_item,omitempty"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
