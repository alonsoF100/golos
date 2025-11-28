package service

import (
	"time"

	apperrors "github.com/alonsoF100/golos/internal/erorrs"
	"github.com/alonsoF100/golos/internal/models"
	"github.com/google/uuid"
)

type Repository interface {
	InsertUser(id, nickname, password string, createdAt time.Time, updatedAt time.Time) (*models.User, error)
	GetUsers() ([]*models.User, error)
	GetUser(id string) (*models.User, error)
	UpdateUser(id, nickname, password string, updatedAt time.Time) (*models.User, error)
	DeleteUser(id string) error
	PatchUser(id string, nickname, password *string, updatedAt time.Time) (*models.User, error)
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

	// TODO добавить хеширование пароля

	user, err := s.repository.InsertUser(id, nickname, password, now, now)
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
	user, err := s.repository.UpdateUser(uuid, nickname, password, now)
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

	user, err := s.repository.PatchUser(uuid, nickname, password, now)
	if err != nil {
		return nil, err
	}

	return user, nil
}
