package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

// TODO добавить методы хранилищу
func (r Repository) UserExist(nicname string) (bool, error) {
	// TODO реализовать функциональность
	return false, nil
}
