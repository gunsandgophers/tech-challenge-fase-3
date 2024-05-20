package repositories

import (
	"errors"
	"tech-challenge-fase-1/internal/adapter/driven/infra/database"
	"tech-challenge-fase-1/internal/core/domain/entities"
)

type CustomerRepositoryDB struct {
	conn database.ConnectionDB
}

func NewCustomerRepositoryDB(conn database.ConnectionDB) *CustomerRepositoryDB {
	return &CustomerRepositoryDB{conn: conn}
}

func (r *CustomerRepositoryDB) GetCustomerByCPF(cpf string) (*entities.Customer, error) {
	return nil, errors.New("NotFound")
}

func (r *CustomerRepositoryDB) Insert(customer *entities.Customer) error {
	sql := "INSERT INTO customer(id, name, email, cpf) VALUES($1, $2, $3, $4)";
	return r.conn.Exec(
		sql,
		customer.GetId(),
		customer.GetName(),
		customer.GetEmail().Value(),
		customer.GetCPF().Value(),
	);
}
