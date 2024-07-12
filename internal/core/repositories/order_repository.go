package repositories

import "tech-challenge-fase-1/internal/core/entities"

type OrderRepositoryInterface interface {
	// Open(customerID *string) (*string, error)
	Insert(order *entities.Order) error
	FindByID(orderID string) (*entities.Order, error)
	// AddOrderItem(order *entities.Order) error
	// UpdateStatus(order *entities.Order) error
}
