package service

import (
	"time"

	"github.com/alonsoF100/golos/internal/models"
	"github.com/alonsoF100/golos/internal/repository/database/postgres"
)

type UserRepository interface {
	CreateUser(id, nickname, password string, createdAt time.Time, updatedAt time.Time) (*models.User, error)
	GetUsers(limit, offset int) ([]*models.User, error)
	GetUser(id string) (*models.User, error)
	GetUserByNickname(nickname string) (*models.User, error)
	UpdateUser(id, nickname, password string, updatedAt time.Time) (*models.User, error)
	DeleteUser(id string) error
	PatchUser(id string, nickname, password *string, updatedAt time.Time) (*models.User, error)
}

type ElectionRepository interface {
	CreateElection(id, userID, name string, description string, createdAt time.Time, updatedAt time.Time) (*models.Election, error)
	GetElections(limit, offset int, userID string) ([]*models.Election, error)
	GetElection(id string) (*models.Election, error)
	DeleteElection(id string) error
	PatchElection(id string, userID, name, description *string, updatedAt time.Time) (*models.Election, error)
}

type VoteVariantRepository interface {
	CreateVoteVariant(id, electionID, name string, createdAt time.Time, updatedAt time.Time) (*models.VoteVariant, error)
	GetVoteVariants(electionID string) ([]*models.VoteVariant, error)
	GetVoteVariant(id string) (*models.VoteVariant, error)
	DeleteVoteVariant(id string) error
	UpdateVoteVariant(id, name string, updatedAt time.Time) (*models.VoteVariant, error)
}

type VoteRepository interface {
	CreateVote(uuid, userID, voteVariantID string, createdAt time.Time, updatedAt time.Time) (*models.Vote, error)
	GetVote(uuid string) (*models.Vote, error)
	GetUserVotes(userID string, voteVariantsIDs []string, limit, offset int) ([]*models.Vote, error)
	GetVariantVotes(voteVariantID string) ([]*models.Vote, error)
	DeleteVote(uuid string) error
	PatchVote(uuid string, userID, voteVariantID *string, updatedAt time.Time) (*models.Vote, error)
}

type UserService struct {
	userRepository UserRepository
}

func NewUser(repository *postgres.Repository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

type ElectionService struct {
	electionRepository ElectionRepository
}

func NewElection(repository *postgres.Repository) *ElectionService {
	return &ElectionService{
		electionRepository: repository,
	}
}

type VoteVariantService struct {
	voteVariantRepository VoteVariantRepository
}

func NewVoteVariant(repository *postgres.Repository) *VoteVariantService {
	return &VoteVariantService{
		voteVariantRepository: repository,
	}
}

type VoteService struct {
	voteRepository VoteRepository
}

func NewVote(repository *postgres.Repository) *VoteService {
	return &VoteService{
		voteRepository: repository,
	}
}

type Service struct {
	*UserService
	*ElectionService
	*VoteVariantService
	*VoteService
}

func New(userRepo, electionRepo, voteVariantRepo, voteRepo *postgres.Repository) *Service {
	return &Service{
		UserService:        NewUser(userRepo),
		ElectionService:    NewElection(electionRepo),
		VoteVariantService: NewVoteVariant(voteVariantRepo),
		VoteService:        NewVote(voteRepo),
	}
}
