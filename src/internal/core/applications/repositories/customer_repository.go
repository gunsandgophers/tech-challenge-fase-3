package repositories

import "tech-challenge-fase-1/internal/core/domain/entities"

type CustomerRepositoryInterface interface {
	GetCustomerByCPF(cpf string) (*entities.Customer, error)
	Insert(*entities.Customer) error
}
