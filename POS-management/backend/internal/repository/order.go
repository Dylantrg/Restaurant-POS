package repository

import "pos-management/internal/models"

type OrderRepository interface {
	Create(o *models.Order) (int, error)
	AddItem(orderID int, item *models.OrderItem) (int, error)
	GetByID(id int) (*models.Order, error)
	UpdateStatus(id int, status string) error
}
