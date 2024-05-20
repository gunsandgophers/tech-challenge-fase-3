package usecases

import (
	"tech-challenge-fase-1/internal/core/applications/repositories"
	"tech-challenge-fase-1/internal/core/domain/dtos"
)

type InsertCustomer struct {
	customerRepository repositories.CustomerRepositoryInterface
}

func NewInsertCustomer(customerRepository repositories.CustomerRepositoryInterface) *InsertCustomer {
	return &InsertCustomer{
		customerRepository: customerRepository,
	}
}

func (ic *InsertCustomer) Execute(dto *dtos.CreateCustomerDTO) (*dtos.CustomerDTO, error) {
	customer, err := dto.ToEntity()
	if err != nil {
		return nil, err
	}

	if err := ic.customerRepository.Insert(customer); err != nil {
		return nil, err
	}
	return dtos.NewCustomerDTOFromEntity(customer), nil
}
