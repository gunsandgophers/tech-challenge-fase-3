package entities

import (
	valueobjects "tech-challenge-fase-1/internal/core/domain/value_objects"

	uuid "github.com/google/uuid"
)

type Customer struct {
	id    string
	name  string
	email string
	cpf   *valueobjects.CPF
}

func CreateCustomer(name string, email string, cpf *valueobjects.CPF) *Customer {
	return &Customer{
		id:    uuid.NewString(),
		name:  name,
		email: email,
		cpf:   cpf,
	}
}

func RestoreCustomer(id string, name string, email string, cpf *valueobjects.CPF) *Customer {
	return &Customer{
		id:    id,
		name:  name,
		email: email,
		cpf:   cpf,
	}
}

func (c *Customer) GetId() string {
	return c.id
}

func (c *Customer) GetName() string {
	return c.name
}

func (c *Customer) GetEmail() string {
	return c.email
}

func (c *Customer) GetCPF() *valueobjects.CPF {
	return c.cpf
}
