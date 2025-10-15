package models

import "time"

// Table represents a restaurant table
type Table struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TableNumber string    `gorm:"unique;not null" json:"table_number"`
	Capacity    int       `gorm:"not null" json:"capacity"`
	QRCodeURL   string    `json:"qr_code_url"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relationships
	Orders []Order `gorm:"foreignKey:TableID" json:"orders,omitempty"`
}

func (Table) TableName() string {
	return "tables"
}
