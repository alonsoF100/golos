package service

import (
	"time"

	apperrors "github.com/alonsoF100/golos/internal/erorrs"
	"github.com/alonsoF100/golos/internal/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s UserService) CreateUser(nickname, password string) (*models.User, error) {
	id := uuid.New().String()
	now := time.Now()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, apperrors.ErrFailedToHashPassword
	}

	user, err := s.userRepository.CreateUser(id, nickname, string(hashedPassword), now, now)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s UserService) GetUsers(limit, offset int) ([]*models.User, error) {
	validLimit := validateLimit(limit)
	validOffset := validateOffset(offset)

	users, err := s.userRepository.GetUsers(validLimit, validOffset)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s UserService) GetUser(uuid string) (*models.User, error) {
	user, err := s.userRepository.GetUser(uuid)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s UserService) UpdateUser(uuid, nickname, password string) (*models.User, error) {
	now := time.Now()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, apperrors.ErrFailedToHashPassword
	}

	user, err := s.userRepository.UpdateUser(uuid, nickname, string(hashedPassword), now)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s UserService) DeleteUser(uuid string) error {
	err := s.userRepository.DeleteUser(uuid)
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) PatchUser(uuid string, nickname, password *string) (*models.User, error) {
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

	user, err := s.userRepository.PatchUser(uuid, nickname, hashedPassword, now)
	if err != nil {
		return nil, err
	}

	return user, nil
}
