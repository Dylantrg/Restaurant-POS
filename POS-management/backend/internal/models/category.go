package models

import "time"

// Category represents a menu category
type Category struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	Description  string    `json:"description"`
	DisplayOrder int       `gorm:"default:0" json:"display_order"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relationships
	MenuItems []MenuItem `gorm:"foreignKey:CategoryID" json:"menu_items,omitempty"`
}

func (Category) TableName() string {
	return "categories"
}
