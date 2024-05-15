package entities

import valueobjects "tech-challenge-fase-1/internal/core/domain/value_objects"

type User struct {
	name string
	email string
	cpf *valueobjects.CPF
}

func NewUser(name string, email string, cpf *valueobjects.CPF) *User {
	return &User{
		name: name,
		email: email,
		cpf: cpf,
	}
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetCPF() *valueobjects.CPF {
	return u.cpf
}
