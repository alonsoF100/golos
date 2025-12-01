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
	err := r.pool.QueryRow(context.Background(), query, id).Scan(
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

// TODO потом тут высрем quary params для сортировки
func (r Repository) GetElections() ([]*models.Election, error) {
	pp := "internal/database/postgres/repository/GetElections"

	const query = `
	SELECT id, user_id, name, description, created_at, updated_at
	FROM elections`

	rows, err := r.pool.Query(context.Background(), query)
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

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////
// TODO реализовать функциональность
