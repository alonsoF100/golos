package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
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

	return pool, nil
}
