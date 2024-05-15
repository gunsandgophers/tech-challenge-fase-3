package repositories

import "tech-challenge-fase-1/internal/core/domain/entities"

type UserRepositoryInterface interface {
	ListUsers() []*entities.User
}
