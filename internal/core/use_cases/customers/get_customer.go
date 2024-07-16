package customers

import (
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/dtos"
	valueobjects "tech-challenge-fase-1/internal/core/value_objects"
)

type GetCustomer struct {
	customerRepository repositories.CustomerRepositoryInterface
}

func NewGetCustomerUseCase(customerRepository repositories.CustomerRepositoryInterface) *GetCustomer {
	return &GetCustomer{
		customerRepository: customerRepository,
	}
}

func (gc *GetCustomer) Execute(cpf string) (*dtos.CustomerDTO, error) {
	cpfVO, err := valueobjects.NewCPF(cpf)
	if err != nil {
		return nil, err
	}
	customer, err := gc.customerRepository.GetCustomerByCPF(cpfVO)
	if err != nil {
		return nil, err
	}
	return dtos.NewCustomerDTOFromEntity(customer), nil
}
