package service

import (
	"time"

	"github.com/alonsoF100/golos/internal/models"
)

type Repository interface {
	CreateUser(id, nickname, password string, createdAt time.Time, updatedAt time.Time) (*models.User, error)
	GetUsers() ([]*models.User, error)
	GetUser(id string) (*models.User, error)
	UpdateUser(id, nickname, password string, updatedAt time.Time) (*models.User, error)
	DeleteUser(id string) error
	PatchUser(id string, nickname, password *string, updatedAt time.Time) (*models.User, error)

	CreateElection(id, userID, name string, description string, createdAt time.Time, updatedAt time.Time) (*models.Election, error)
	GetElections() ([]*models.Election, error)
	GetElection(id string) (*models.Election, error)
	DeleteElection(id string) error
	PatchElection(id string, userID, name, description *string, updatedAt time.Time) (*models.Election, error)

	CreateVoteVariant(id, electionID, name string, createdAt time.Time, updatedAt time.Time) (*models.VoteVariant, error)
	GetVoteVariants(electionID string) ([]*models.VoteVariant, error)
	GetVoteVariant(id string) (*models.VoteVariant, error)
	DeleteVoteVariant(id string) error
	UpdateVoteVariant(id, name string, updatedAt time.Time) (*models.VoteVariant, error)
}

type Service struct {
	repository Repository
}

func New(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}