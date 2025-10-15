package models

import "time"

// MenuItem represents a menu item
type MenuItem struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	CategoryID    uint      `gorm:"not null" json:"category_id"`
	Name          string    `gorm:"not null" json:"name"`
	Description   string    `json:"description"`
	Price         float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	CostPrice     float64   `gorm:"type:decimal(10,2)" json:"cost_price"`
	StockQuantity int       `gorm:"default:0" json:"stock_quantity"`
	IsAvailable   bool      `gorm:"default:true" json:"is_available"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relationships
	Category          Category           `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	OrderItems        []OrderItem        `gorm:"foreignKey:MenuItemID" json:"order_items,omitempty"`
	StockTransactions []StockTransaction `gorm:"foreignKey:MenuItemID" json:"stock_transactions,omitempty"`
}

func (MenuItem) TableName() string {
	return "menu_items"
}
