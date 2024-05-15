package infra

import (
	"tech-challenge-fase-1/internal/core/domain/entities"
	valueobjects "tech-challenge-fase-1/internal/core/domain/value_objects"
)

type UserRepositoryDB struct {
	/*conn*/
}

func NewUserRepositoryDB(/*conn*/) *UserRepositoryDB {
	return &UserRepositoryDB{}
}

func (r *UserRepositoryDB) ListUsers() []*entities.User {
	var users []*entities.User
	users = append(
		users,
		entities.NewUser(
			"Gabriel",
			"gabriel@mail.com",
			valueobjects.NewCPF("000.000.000-00"),
		),
	)
	users = append(
		users,
		entities.NewUser(
			"Rafael",
			"rafael@mail.com",
			valueobjects.NewCPF("111.111.111-11"),
		),
	)
	users = append(
		users,
		entities.NewUser(
			"Nath",
			"nath@mail.com",
			valueobjects.NewCPF("222.222.222-22"),
		),
	)
	users = append(
		users,
		entities.NewUser(
			"Samuel",
			"samuel@mail.com",
			valueobjects.NewCPF("333.333.333-33"),
		),
	)
	return users
}
