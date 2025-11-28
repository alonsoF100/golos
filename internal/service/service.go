package service

import (
	"fmt"
	"time"

	apperrors "github.com/alonsoF100/golos/internal/erorrs"
	"github.com/alonsoF100/golos/internal/models"
	"github.com/google/uuid"
)

type Repository interface {
	// TODO добавить контракты для интерфейса
	UserExist(nickname string) (bool, error)
	InsertUser(id, nickname, password string, createdAt time.Time, updatedAt time.Time) (*models.User, error)
}

type Service struct {
	repository Repository
}

func New(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// TODO выдавить документацию к функциям малую

func (s Service) CreateUser(nickname, password string) (*models.User, error) {
	exist, err := s.repository.UserExist(nickname)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, apperrors.ErrUserAlreadyExist
	}

	// TODO добавить хеширование пароля
	t := time.Now()
	id := uuid.New().String()
	createdAt := t
	updatedAt := t

	user, err := s.repository.InsertUser(id, nickname, password, createdAt, updatedAt)
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
