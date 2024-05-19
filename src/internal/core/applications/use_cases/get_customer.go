package usecases

import (
	"tech-challenge-fase-1/internal/core/applications/repositories"
	"tech-challenge-fase-1/internal/core/domain/dtos"
)

type GetCustomer struct {
	customerRepository repositories.CustomerRepositoryInterface
}

func NewGetCustomer(customerRepository repositories.CustomerRepositoryInterface) *GetCustomer {
	return &GetCustomer{
		customerRepository: customerRepository,
	}
}

func (gc *GetCustomer) Execute(cpf string) (*dtos.CustomerDTO, error) {
	customers, err := gc.customerRepository.GetCustomerByCPF(cpf)
	if err != nil {
		return nil, err
	}
	return dtos.NewCustomerDTOFromEntity(customers), nil
}
