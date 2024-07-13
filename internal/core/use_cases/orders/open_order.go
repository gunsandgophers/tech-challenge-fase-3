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

func (co *OpenOrderUseCase) validateCustomerId(customerId *string) error {
	if customerId == nil {
		return nil
	}
	_, err := co.customerRepository.GetCustomerByID(*customerId)
	if err != nil {
		return err
	}
	return nil
}

func (co *OpenOrderUseCase) Execute(customerId *string) (*dtos.OrderDTO, error) {
	if err := co.validateCustomerId(customerId); err != nil {
		return nil, err
	}
	order := entities.CreateOpenOrder(customerId)
	if err := co.orderRepository.Insert(order); err != nil {
		return nil, err
	}
	return dtos.NewOrderDTOFromEntity(order), nil
}
