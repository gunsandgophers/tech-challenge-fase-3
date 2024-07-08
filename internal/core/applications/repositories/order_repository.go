package repositories

import "tech-challenge-fase-1/internal/core/domain/entities"

type OrderRepositoryInterface interface {
	Open(customerID *string) (*string, error)
	FindByID(orderID *string) (*entities.Order, error)
	AddOrderItem(order *entities.Order) error
	UpdateStatus(order *entities.Order) error
}
