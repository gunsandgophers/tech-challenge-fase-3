package usecases

import (
	"tech-challenge-fase-1/internal/core/applications/repositories"
	"tech-challenge-fase-1/internal/core/domain/dtos"
)

type ListUsers struct {
	userRepository repositories.UserRepositoryInterface
}

func NewListUsers(userRepository repositories.UserRepositoryInterface) *ListUsers {
	return &ListUsers{
		userRepository: userRepository,
	}
}

func (lu *ListUsers) Execute() []dtos.UserDTO {
	users := lu.userRepository.ListUsers()
	var usersDtos []dtos.UserDTO
	for _, user := range users {
		usersDtos = append(usersDtos, *dtos.NewUserDTOFromEntity(user));
	}
	return usersDtos
}
