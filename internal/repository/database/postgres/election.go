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
)

func (r Repository) CreateElection(id, userID, name string, description string, updatedAt time.Time, createdAt time.Time) (*models.Election, error) {
	pp := "internal/database/postgres/repository/CreateElection"

	const query = `
	INSERT INTO elections (id, user_id, name, description, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, user_id, name, description, created_at, updated_at;`

	var election models.Election
	err := r.pool.QueryRow(context.Background(), query, id, userID, name, description, createdAt, updatedAt).Scan(
		&election.ID,
		&election.UserID,
		&election.Name,
		&election.Description,
		&election.CreatedAt,
		&election.UpdatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return nil, apperrors.ErrUserNotFound
		}
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	return &election, nil
}

func (r Repository) GetElections(limit, offset int, userID string) ([]*models.Election, error) {
	pp := "internal/database/postgres/repository/GetElections"

	qb := squirrel.
		Select("id", "user_id", "name", "description", "created_at", "updated_at").
		From("elections")
	if userID != "" {
		qb = qb.Where(squirrel.Eq{"user_id": userID})
	}
	query, args, err := qb.
		OrderBy("created_at DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	rows, err := r.pool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}
	defer rows.Close()

	var elections []*models.Election
	for rows.Next() {
		var election models.Election
		err := rows.Scan(
			&election.ID,
			&election.UserID,
			&election.Name,
			&election.Description,
			&election.CreatedAt,
			&election.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: error: %w", pp, err)
		}

		elections = append(elections, &election)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	return elections, nil
}

func (r Repository) GetElection(id string) (*models.Election, error) {
	pp := "internal/database/postgres/repository/GetElection"

	const query = `
	SELECT id, user_id, name, description, created_at, updated_at
	FROM elections
	WHERE id = $1`

	var election models.Election
	err := r.pool.QueryRow(context.Background(), query, id).Scan(
		&election.ID,
		&election.UserID,
		&election.Name,
		&election.Description,
		&election.CreatedAt,
		&election.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrElectionNotFound
		}
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	return &election, nil
}

func (r Repository) DeleteElection(id string) error {
	pp := "internal/database/postgres/repository/DeleteElection"

	const query = `
	DELETE FROM elections
	WHERE id = $1`

	row, err := r.pool.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("%s: error: %w", pp, err)
	}

	if row.RowsAffected() == 0 {
		return apperrors.ErrElectionNotFound
	}

	return nil
}

func (r Repository) PatchElection(id string, userID, name, description *string, updatedAt time.Time) (*models.Election, error) {
	pp := "internal/database/postgres/repository/PatchElection"

	qb := squirrel.Update("elections").
		Set("updated_at", updatedAt).
		Where(squirrel.Eq{"id": id})
	if userID != nil {
		qb = qb.Set("user_id", *userID)
	}
	if name != nil {
		qb = qb.Set("name", *name)
	}
	if description != nil {
		qb = qb.Set("description", *description)
	}
	query, args, err := qb.
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id, user_id, name, description, created_at, updated_at").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	var election models.Election
	err = r.pool.QueryRow(context.Background(), query, args...).Scan(
		&election.ID,
		&election.UserID,
		&election.Name,
		&election.Description,
		&election.CreatedAt,
		&election.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrElectionNotFound
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return nil, apperrors.ErrUserNotFound
		}
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	return &election, nil
}
