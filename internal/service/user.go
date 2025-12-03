package service

import (
	"time"

	apperrors "github.com/alonsoF100/golos/internal/erorrs"
	"github.com/alonsoF100/golos/internal/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

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
