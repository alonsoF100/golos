package service

import (
	"time"

	"github.com/alonsoF100/golos/internal/models"
)

const (
	maxLimit     = 100
	defaultLimit = 20
)

type Repository interface {
	// User repository methods
	CreateUser(id, nickname, password string, createdAt time.Time, updatedAt time.Time) (*models.User, error)
	GetUsers(limit, offset int) ([]*models.User, error)
	GetUser(id string) (*models.User, error)
	GetUserByNickname(nickname string) (*models.User, error)
	UpdateUser(id, nickname, password string, updatedAt time.Time) (*models.User, error)
	DeleteUser(id string) error
	PatchUser(id string, nickname, password *string, updatedAt time.Time) (*models.User, error)

	// Election repository methods
	CreateElection(id, userID, name string, description string, createdAt time.Time, updatedAt time.Time) (*models.Election, error)
	GetElections(limit, offset int, userID string) ([]*models.Election, error)
	GetElection(id string) (*models.Election, error)
	DeleteElection(id string) error
	PatchElection(id string, userID, name, description *string, updatedAt time.Time) (*models.Election, error)

	// VoteVariant repository methods
	CreateVoteVariant(id, electionID, name string, createdAt time.Time, updatedAt time.Time) (*models.VoteVariant, error)
	GetVoteVariants(electionID string) ([]*models.VoteVariant, error)
	GetVoteVariant(id string) (*models.VoteVariant, error)
	DeleteVoteVariant(id string) error
	UpdateVoteVariant(id, name string, updatedAt time.Time) (*models.VoteVariant, error)

	// Vote repository methods
	CreateVote(uuid, userID, voteVariantID string, createdAt time.Time, updatedAt time.Time) (*models.Vote, error)
	GetVote(uuid string) (*models.Vote, error)
	GetUserVotes(userID string) ([]*models.Vote, error)
	GetVariantVotes(voteVariantID string) ([]*models.Vote, error)
	DeleteVote(uuid string) error
	PatchVote(uuid string, userID, voteVariantID *string, updatedAt time.Time) (*models.Vote, error)
}

type Service struct {
	repository Repository
}

func New(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func validateLimit(limit int) int {
	if limit <= 0 {
		return defaultLimit
	}
	if limit > maxLimit {
		return maxLimit
	}
	return limit
}

func validateOffset(offset int) int {
	if offset < 0 {
		return 0
	}

	return offset
}
