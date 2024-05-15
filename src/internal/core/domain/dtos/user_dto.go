package dtos

import "tech-challenge-fase-1/internal/core/domain/entities"

type UserDTO struct {
	Name string
	Email string
	Cpf string
}

func NewUserDTOFromEntity(user *entities.User) *UserDTO {
	return &UserDTO{
		Name: user.GetName(),
		Email: user.GetEmail(),
		Cpf: user.GetCPF().Formated(),
	}
}
