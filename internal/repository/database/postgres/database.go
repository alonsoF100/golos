package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func NewPool(connString string) (*pgxpool.Pool, error) {

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	// TODO добавлять необходимые настройки pool а

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	// миграции
	connConfig := poolConfig.ConnConfig
	db := stdlib.OpenDB(*connConfig)

	if err := goose.Up(db, "./migrations/postgres"); err != nil {
		return nil, err
	}

	return pool, nil
}
