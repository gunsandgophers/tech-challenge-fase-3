package orders

import (
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/core/repositories"
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

func (co *OpenOrderUseCase) Execute(customerId string) (*dtos.OrderDTO, error) {
	if customerId != "" {
		_ , err := co.customerRepository.GetCustomerByID(customerId)
		if err != nil {
			return nil, err
		}
	}
	order := entities.CreateOpenOrder(customerId)
	err := co.orderRepository.Insert(order)
	if err != nil {
		return nil, err
	}
	return dtos.NewOrderDTOFromEntity(order), nil
}
