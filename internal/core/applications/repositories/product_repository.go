package repositories

import "tech-challenge-fase-1/internal/core/domain/entities"

type ProductRepositoryInterface interface {
	CreateProduct(product *entities.Product) error
	FindByID(ID int) (*entities.Product, error)
	UpdateProduct(product *entities.Product) (*entities.Product, error)
	DeleteProduct(ID int) error
	FindByCategory(category string, page, size int) ([]*entities.Product, error)
	ListProducts() ([]*entities.Product, error)
}
