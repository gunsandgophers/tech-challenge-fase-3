package orders

import (
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/entities"
)

type OpenOrderUseCase struct {
	orderRepository    repositories.OrderRepositoryInterface
	customerRepository repositories.CustomerRepositoryInterface
}

func NewOpenOrderUseCase(
	orderRepository repositories.OrderRepositoryInterface,
	customerRepository repositories.CustomerRepositoryInterface,
) *OpenOrderUseCase {
	return &OpenOrderUseCase{
		orderRepository:    orderRepository,
		customerRepository: customerRepository,
	}
}

func (co *OpenOrderUseCase) Execute(customerID *string) (*dtos.OrderDTO, error) {

	if customerID != nil {
		_, err := co.customerRepository.GetCustomerByID(customerID)

		if err != nil {
			return nil, err
		}
	}

	id, err := co.orderRepository.Open(customerID)

	if err != nil {
		return nil, err
	}

	return dtos.NewOpenOrderDTOFromEntity(&entities.Order{
		ID: *id,
	}), nil
}
