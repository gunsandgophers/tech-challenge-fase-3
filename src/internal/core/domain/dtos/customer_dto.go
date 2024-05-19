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
		Name: customer.GetName(),
		Email: customer.GetEmail(),
		Cpf: customer.GetCPF().Formated(),
	}
}
