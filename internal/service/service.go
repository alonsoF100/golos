package service

import (
	"fmt"
	"time"

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
	// TODO высрать функциональность

	// Пока загшлука на вывод, чтобы красным не гадило
	return nil, nil
}

func (s Service) GetUser(uuid string) (*models.User, error) {
	// TODO высрать функциональность

	// Пока загшлука на вывод, чтобы красным не гадило
	return nil, nil
}

func (s Service) UpdateUser(uuid, nickname, password string) (*models.User, error) {
	// TODO высрать функциональность

	// Пока загшлука на вывод, чтобы красным не гадило
	return nil, nil
}

func (s Service) DeleteUser(uuid string) error {
	// TODO высрать функциональность

	// Пока загшлука на вывод, чтобы красным не гадило
	return nil
}

func (s Service) PatchUser(uuid string, nickname, password *string) (*models.User, error) {
	pp := "internal/service/PathUser"
	// TODO высрать функциональность
	if nickname == nil && password == nil {
		return nil, fmt.Errorf("%s: error: no fields to update", pp)
	}

	// Пока загшлука на вывод, чтобы красным не гадило
	return nil, nil
}
