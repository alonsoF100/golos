package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	apperrors "github.com/alonsoF100/golos/internal/erorrs"
	"github.com/alonsoF100/golos/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////
// TODO реализовать функциональность
func (r Repository) CreateVoteVariant(id, electionID, name string, createdAt time.Time, updatedAt time.Time) (*models.VoteVariant, error) {
	pp := "internal/database/postgres/repository/CreateVoteVariant"

	const query = `
	INSERT INTO vote_variants (id, election_id, name, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, election_id, name, created_at, updated_at`

	var voteVariant models.VoteVariant
	err := r.pool.QueryRow(context.Background(), query, id, electionID, name, createdAt, updatedAt).Scan(
		&voteVariant.ID,
		&voteVariant.ElectionID,
		&voteVariant.Name,
		&voteVariant.CreatedAt,
		&voteVariant.UpdatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return nil, apperrors.ErrElectionNotFound
		}
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	return &voteVariant, nil
}

// TODO потом тут высрем quary params для сортировки
func (r Repository) GetVoteVariants(electionID string) ([]*models.VoteVariant, error) {
	pp := "internal/database/postgres/repository/GetVoteVariants"

	const query = `
	SELECT id, election_id, name, created_at, updated_at FROM vote_variants
	WHERE election_id = $1`

	rows, err := r.pool.Query(context.Background(), query, electionID)
	if err != nil {
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}
	defer rows.Close()

	var voteVariants []*models.VoteVariant
	for rows.Next() {
		var voteVariant models.VoteVariant
		err := rows.Scan(
			&voteVariant.ID,
			&voteVariant.ElectionID,
			&voteVariant.Name,
			&voteVariant.CreatedAt,
			&voteVariant.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: error: %w", pp, err)
		}

		voteVariants = append(voteVariants, &voteVariant)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	return voteVariants, nil
}

func (r Repository) GetVoteVariant(id string) (*models.VoteVariant, error) {
	pp := "internal/database/postgres/repository/GetVoteVariant"

	const query = `
	SELECT id, election_id, name, created_at, updated_at FROM vote_variants
	WHERE id = $1`

	var voteVariant models.VoteVariant
	err := r.pool.QueryRow(context.Background(), query, id).Scan(
		&voteVariant.ID,
		&voteVariant.ElectionID,
		&voteVariant.Name,
		&voteVariant.CreatedAt,
		&voteVariant.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrVoteVariantNotFound
		}
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	return &voteVariant, nil
}

func (r Repository) DeleteVoteVariant(id string) error {
	pp := "internal/database/postgres/repository/DeleteVoteVariant"

	const query = `
	DELETE FROM vote_variants
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

func (r Repository) UpdateVoteVariant(id, name string, updatedAt time.Time) (*models.VoteVariant, error) {
	pp := "internal/database/postgres/repository/UpdateVoteVariant"

	const query = `
	UPDATE vote_variants
	SET name = $1, updated_at = $2
	WHERE id = $3
	RETURNING id, election_id, name, created_at, updated_at`

	var voteVariant models.VoteVariant
	err := r.pool.QueryRow(context.Background(), query, name, updatedAt, id).Scan(
		&voteVariant.ID,
		&voteVariant.ElectionID,
		&voteVariant.Name,
		&voteVariant.CreatedAt,
		&voteVariant.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrVoteVariantNotFound
		}
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	return &voteVariant, nil
}
