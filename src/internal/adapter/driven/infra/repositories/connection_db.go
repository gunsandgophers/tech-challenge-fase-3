package infra

import (
	"context"
	"fmt"
	"log"
	"tech-challenge-fase-1/internal/adapter/driven/infra/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConnection() *pgxpool.Pool {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DB_HOST, config.DB_POST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME)

	pool, err := pgxpool.New(context.Background(), psqlInfo)
	if err != nil {
		log.Println("Unable to create connection pool", err)
		panic(1)
	}

	return pool
}
