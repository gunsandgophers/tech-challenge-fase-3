package dtos

import "tech-challenge-fase-1/internal/core/domain/entities"

type CustomerDTO struct {
	Id string
	Name string
	Email string
	Cpf string
}

func NewCustomerDTOFromEntity(customer *entities.Customer) *CustomerDTO {
	return &CustomerDTO{
		Id: customer.GetId(),
		Name: customer.GetName(),
		Email: customer.GetEmail().Value(),
		Cpf: customer.GetCPF().Value(),
	}
}

type CreateCustomerDTO struct {
	Name string
	Email string
	Cpf string
}

func (cc *CreateCustomerDTO) ToEntity() (*entities.Customer, error) {
	return entities.CreateCustomer(cc.Name, cc.Email, cc.Cpf)
}
