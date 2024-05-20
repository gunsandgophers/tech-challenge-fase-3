package usecases

import (
	"errors"
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
	if existCustomer, _ := ic.customerRepository.GetCustomerByCPF(customer.GetCPF()); existCustomer != nil {
		return nil, errors.New("Existing Customer with this CPF")
	}
	if existCustomer, _ := ic.customerRepository.GetCustomerByEmail(customer.GetEmail()); existCustomer != nil {
		return nil, errors.New("Existing Customer with this Email")
	}

	if err := ic.customerRepository.Insert(customer); err != nil {
		return nil, err
	}
	return dtos.NewCustomerDTOFromEntity(customer), nil
}
