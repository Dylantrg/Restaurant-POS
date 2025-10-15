package repository

import "pos-management/internal/models"

type ProductRepository interface {
	GetByID(id int) (*models.Product, error)
	ListActive() ([]models.Product, error)
	Create(p *models.Product) (int, error)
	Update(p *models.Product) error
	Deactivate(id int) error
}
