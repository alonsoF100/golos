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

func (r Repository) CreateVote(uuid, userID, voteVariantID string, createdAt time.Time, updatedAt time.Time) (*models.Vote, error) {
	pp := "internal/database/postgres/repository/CreateVote"

	const query = `
	INSERT INTO votes (id, user_id, variant_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, user_id, variant_id, created_at, updated_at`

	var vote models.Vote
	err := r.pool.QueryRow(context.Background(), query, uuid, userID, voteVariantID, createdAt, updatedAt).Scan(
		&vote.ID,
		&vote.UserID,
		&vote.VariantID,
		&vote.CreatedAt,
		&vote.UpdatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return nil, errors.Join(apperrors.ErrUserNotFound, apperrors.ErrVoteVariantNotFound)
		}
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	return &vote, nil
}

func (r Repository) GetVote(uuid string) (*models.Vote, error) {
	pp := "internal/database/postgres/repository/GetVote"

	const query = `
	SELECT id, user_id, variant_id, created_at, updated_at FROM votes
	WHERE id = $1`

	var vote models.Vote
	err := r.pool.QueryRow(context.Background(), query, uuid).Scan(
		&vote.ID,
		&vote.UserID,
		&vote.VariantID,
		&vote.CreatedAt,
		&vote.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrVoteNotFound
		}
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	return &vote, nil
}

func (r Repository) DeleteVote(uuid string) error {
	pp := "internal/database/postgres/repository/DeleteVote"

	const query = `
	DELETE FROM votes
	WHERE id = $1`

	row, err := r.pool.Exec(context.Background(), query, uuid)
	if err != nil {
		return fmt.Errorf("%s: error: %w", pp, err)
	}

	if row.RowsAffected() == 0 {
		return apperrors.ErrVoteNotFound
	}

	return nil
}

func (r Repository) PatchVote(uuid string, userID, voteVariantID *string, updatedAt time.Time) (*models.Vote, error) {
	pp := "internal/database/postgres/repository/PatchVote"

	qb := squirrel.Update("votes").
		Set("updated_at", updatedAt).
		Where(squirrel.Eq{"id": uuid})
	if userID != nil {
		qb = qb.Set("user_id", *userID)
	}
	if voteVariantID != nil {
		qb = qb.Set("variant_id", *voteVariantID)
	}
	query, args, err := qb.
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id, user_id, variant_id, created_ad, updated_at").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	var vote models.Vote
	err = r.pool.QueryRow(context.Background(), query, args...).Scan(
		&vote.ID,
		&vote.UserID,
		&vote.VariantID,
		&vote.CreatedAt,
		&vote.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrVoteNotFound
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return nil, errors.Join(apperrors.ErrUserNotFound, apperrors.ErrVoteVariantNotFound)
		}
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	return &vote, nil
}

func (r Repository) GetUserVotes(userID string, voteVariantsIDs []string, limit, offset int) ([]*models.Vote, error) {
	pp := "internal/database/postgres/repository/GetUserVotes"

	qb := squirrel.Select("id", "user_id", "variant_id", "created_at", "updated_at").
		From("votes").
		Where(squirrel.Eq{"user_id": userID})

	if len(voteVariantsIDs) > 0 {
		qb = qb.Where(squirrel.Eq{"variant_id": voteVariantsIDs})
	}

	query, args, err := qb.OrderBy("created_at DESC").
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

	var votes []*models.Vote
	for rows.Next() {
		var vote models.Vote

		err := rows.Scan(
			&vote.ID,
			&vote.UserID,
			&vote.VariantID,
			&vote.CreatedAt,
			&vote.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: error: %w", pp, err)
		}

		votes = append(votes, &vote)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: error: %w", pp, err)
	}

	return votes, nil
}

func (r Repository) GetVariantVotes(voteVariantID string) ([]*models.Vote, error) {
	return nil, nil
}
