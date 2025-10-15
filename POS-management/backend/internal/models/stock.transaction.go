package models

import "time"

// TransactionType represents the type of stock transaction
type TransactionType string

const (
	TransactionTypeIn         TransactionType = "in"
	TransactionTypeOut        TransactionType = "out"
	TransactionTypeAdjustment TransactionType = "adjustment"
)

// StockTransaction represents a stock movement transaction
type StockTransaction struct {
	ID              uint            `gorm:"primaryKey" json:"id"`
	MenuItemID      uint            `gorm:"not null" json:"menu_item_id"`
	UserID          uint            `gorm:"not null" json:"user_id"`
	TransactionType TransactionType `gorm:"type:varchar(20);not null" json:"transaction_type"`
	QuantityChange  int             `gorm:"not null" json:"quantity_change"`
	Reason          string          `json:"reason"`
	CostPerUnit     float64         `gorm:"type:decimal(10,2)" json:"cost_per_unit"`
	CreatedAt       time.Time       `gorm:"autoCreateTime" json:"created_at"`

	// Relationships
	MenuItem MenuItem `gorm:"foreignKey:MenuItemID" json:"menu_item,omitempty"`
	User     User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (StockTransaction) TableName() string {
	return "stock_transactions"
}
