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

	CreateElection(id, userID, name string, description *string, updatedAt time.Time, createdAt time.Time) (*models.Election, error)
}

type Service struct {
	repository Repository
}

func New(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

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
	
}
