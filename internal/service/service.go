package service

import (
	"time"

	apperrors "github.com/alonsoF100/golos/internal/erorrs"
	"github.com/alonsoF100/golos/internal/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	CreateUser(id, nickname, password string, createdAt time.Time, updatedAt time.Time) (*models.User, error)
	GetUsers() ([]*models.User, error)
	GetUser(id string) (*models.User, error)
	UpdateUser(id, nickname, password string, updatedAt time.Time) (*models.User, error)
	DeleteUser(id string) error
	PatchUser(id string, nickname, password *string, updatedAt time.Time) (*models.User, error)

	CreateElection(id, userID, name string, description *string, createdAt time.Time, updatedAt time.Time) (*models.Election, error)
	GetElections() ([]*models.Election, error)
	GetElection(id string) (*models.Election, error)
	DeleteElection(id string) error
	PatchElection(id string, userID, name, description *string, updatedAt time.Time) (*models.Election, error)

	CreateVoteVariant(id, electionID, name string, createdAt time.Time, updatedAt time.Time) (*models.VoteVariant, error)
	GetVoteVariants() ([]*models.VoteVariant, error)
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

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////
func (s Service) CreateUser(nickname, password string) (*models.User, error) {
	id := uuid.New().String()
	now := time.Now()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, apperrors.ErrFailedToHashPassword
	}

	user, err := s.repository.CreateUser(id, nickname, string(hashedPassword), now, now)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s Service) GetUsers() ([]*models.User, error) {
	users, err := s.repository.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s Service) GetUser(uuid string) (*models.User, error) {
	user, err := s.repository.GetUser(uuid)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s Service) UpdateUser(uuid, nickname, password string) (*models.User, error) {
	now := time.Now()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, apperrors.ErrFailedToHashPassword
	}

	user, err := s.repository.UpdateUser(uuid, nickname, string(hashedPassword), now)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s Service) DeleteUser(uuid string) error {
	err := s.repository.DeleteUser(uuid)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) PatchUser(uuid string, nickname, password *string) (*models.User, error) {
	now := time.Now()
	if nickname == nil && password == nil {
		return nil, apperrors.ErrNothingToChange
	}

	var hashedPassword *string
	if password != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
		if err != nil {
			return nil, apperrors.ErrFailedToHashPassword
		}
		hashedStr := string(hashed)
		hashedPassword = &hashedStr
	}

	user, err := s.repository.PatchUser(uuid, nickname, hashedPassword, now)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////
func (s Service) CreateElection(userID string, name string, description *string) (*models.Election, error) {
	now := time.Now()
	id := uuid.New().String()

	election, err := s.repository.CreateElection(id, userID, name, description, now, now)
	if err != nil {
		return nil, err
	}

	return election, nil
}

func (s Service) GetElections() ([]*models.Election, error) {
	elections, err := s.repository.GetElections()
	if err != nil {
		return nil, err
	}

	return elections, nil
}

func (s Service) GetElection(uuid string) (*models.Election, error) {
	election, err := s.repository.GetElection(uuid)
	if err != nil {
		return nil, err
	}

	return election, nil
}

func (s Service) DeleteElection(uuid string) error {
	err := s.repository.DeleteElection(uuid)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) PatchElection(uuid string, userID, name, description *string) (*models.Election, error) {
	now := time.Now()
	if userID == nil && name == nil && description == nil {
		return nil, apperrors.ErrNothingToChange
	}

	election, err := s.repository.PatchElection(uuid, userID, name, description, now)
	if err != nil {
		return nil, err
	}

	return election, nil
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////
func (s Service) CreateVoteVariant(electionID, name string) (*models.VoteVariant, error) {
	now := time.Now()
	id := uuid.New().String()

	voteVariant, err := s.repository.CreateVoteVariant(id, electionID, name, now, now)
	if err != nil {
		return nil, err
	}

	return voteVariant, nil
}

func (s Service) GetVoteVariants() ([]*models.VoteVariant, error) {
	voteVariants, err := s.repository.GetVoteVariants()
	if err != nil {
		return nil, err
	}

	return voteVariants, nil
}

func (s Service) GetVoteVariant(uuid string) (*models.VoteVariant, error) {
	voteVariant, err := s.repository.GetVoteVariant(uuid)
	if err != nil {
		return nil, err
	}

	return voteVariant, nil
}

func (s Service) DeleteVoteVariant(uuid string) error {
	err := s.repository.DeleteVoteVariant(uuid)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) UpdateVoteVariant(uuid string, name string) (*models.VoteVariant, error) {
	now := time.Now()

	voteVariant, err := s.repository.UpdateVoteVariant(uuid, name, now)
	if err != nil {
		return nil, err
	}

	return voteVariant, nil
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////
//TODO votes потом сделаем