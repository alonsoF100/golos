package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	apperrors "github.com/alonsoF100/golos/internal/erorrs"
	"github.com/alonsoF100/golos/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r Repository) CreateUser(id, nickname, password string, createdAt time.Time, updatedAt time.Time) (*models.User, error) {
	pp := "internal/database/postgres/repository/CreatetUser"

	const query = `
	INSERT INTO users (id, nickname, password, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, nickname, password, created_at, updated_at`

	var user models.User
	err := r.pool.QueryRow(context.Background(), query, id, nickname, password, createdAt, updatedAt).Scan(
		&user.ID,
		&user.Nickname,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, apperrors.ErrUserAlreadyExist
		}
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	return &user, nil
}

// TODO потом тут высрем quary params для сортировки братков
func (r Repository) GetUsers() ([]*models.User, error) {
	pp := "internal/database/postgres/repository/GetUsers"

	const query = `
	SELECT id, nickname, password, created_at, updated_at FROM users`

	rows, err := r.pool.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Nickname,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: error: %w", pp, err)
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	return users, nil
}

func (r Repository) GetUser(id string) (*models.User, error) {
	pp := "internal/database/postgres/repository/GetUser"

	const query = `
	SELECT id, nickname, password, created_at, updated_at FROM users
	WHERE id = $1`

	var user models.User
	err := r.pool.QueryRow(context.Background(), query, id).Scan(&user)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrUserNotFound
		}
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	return &user, nil
}

func (r Repository) UpdateUser(id, nickname, password string, updatedAt time.Time) (*models.User, error) {
	pp := "internal/database/postgres/repository/UpdateUser"

	const query = `
	UPDATE users
	SET nickname = $1, password = $2, updated_at = $3
	WHERE id = $4
	RETURNING id, nickname, password, created_at, updated_at`

	var user models.User
	err := r.pool.QueryRow(context.Background(), query, nickname, password, updatedAt, id).Scan(
		&user.ID,
		&user.Nickname,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrUserNotFound
		}
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	return &user, nil
}

func (r Repository) DeleteUser(id string) error {
	pp := "internal/database/postgres/repository/DeleteUser"

	const query = `
	DELETE FROM users
	WHERE id = $1`

	row, err := r.pool.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("%s: error: %w", pp, err)
	}

	if row.RowsAffected() == 0 {
		return apperrors.ErrUserNotFound
	}

	return nil
}

func (r Repository) PatchUser(id string, nickname, password *string, updatedAt time.Time) (*models.User, error) {
	pp := "internal/database/postgres/repository/PatchUser"

	qb := squirrel.Update("users").
		Set("updated_at", updatedAt).
		Where(squirrel.Eq{"id": id})
	if nickname != nil {
		qb = qb.Set("nickname", *nickname)
	}
	if password != nil {
		qb = qb.Set("password", *password)
	}
	query, args, err := qb.
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id, nickname, password, created_at, updated_at").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	var user models.User
	err = r.pool.QueryRow(context.Background(), query, args...).Scan(
		&user.ID,
		&user.Nickname,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrUserNotFound
		}
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	return &user, nil
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////
// TODO реализовать функциональность
func (r Repository) CreateElection(id, userID, name string, description *string, updatedAt time.Time, createdAt time.Time) (*models.Election, error) {
	return nil, nil
}

func (r Repository) GetElections() ([]*models.Election, error) {
	return nil, nil
}

func (r Repository) GetElection(id string) (*models.Election, error) {
	return nil, nil
}

func (r Repository) DeleteElection(id string) error {
	return nil
}

func (r Repository) PatchElection(id string, userID, name, description *string, updatedAt time.Time) (*models.Election, error) {
	return nil, nil
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////
// TODO реализовать функциональность
func (r Repository) CreateVoteVariant(id, electionID, name string, createdAt time.Time, updatedAt time.Time) (*models.VoteVariant, error) {
	return nil, nil
}

func (r Repository) GetVoteVariants() ([]*models.VoteVariant, error) {
	return nil, nil
}

func (r Repository) GetVoteVariant(id string) (*models.VoteVariant, error) {
	return nil, nil
}

func (r Repository) DeleteVoteVariant(id string) error {
	return nil
}

func (r Repository) UpdateVoteVariant(id, name string, updatedAt time.Time) (*models.VoteVariant, error) {
	return nil, nil
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////
// TODO реализовать функциональность
