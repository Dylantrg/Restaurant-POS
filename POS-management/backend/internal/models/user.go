package models

import (
	"time"
)

// UserRole represents the role of a user in the system
type UserRole string

const (
	RoleWaiter  UserRole = "waiter"
	RoleKitchen UserRole = "kitchen"
	RoleAdmin   UserRole = "admin"
)

// User represents a system user
type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"unique;not null" json:"username"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Email        string    `gorm:"not null" json:"email"`
	Role         UserRole  `gorm:"type:varchar(20);not null" json:"role"`
	Name         string    `gorm:"not null" json:"name"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relationships
	Orders            []Order            `gorm:"foreignKey:WaiterID" json:"orders,omitempty"`
	StockTransactions []StockTransaction `gorm:"foreignKey:UserID" json:"stock_transactions,omitempty"`
}

func (User) TableName() string {
	return "users"
}
