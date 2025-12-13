package service

import (
	"time"

	apperrors "github.com/alonsoF100/golos/internal/erorrs"
	"github.com/alonsoF100/golos/internal/models"
	"github.com/google/uuid"
)

func (s Service) CreateElection(userID string, name string, description string) (*models.Election, error) {
	now := time.Now()
	id := uuid.New().String()

	election, err := s.repository.CreateElection(id, userID, name, description, now, now)
	if err != nil {
		return nil, err
	}

	return election, nil
}

func (s Service) GetElections(limit, offset int, nickname string) ([]*models.Election, error) {
	validateLimit := validateLimit(limit)
	validateOffset := validateOffset(offset)

	user, err := s.repository.GetUserByNickname(nickname)
	if err != nil {
		return nil, apperrors.ErrUserNotFound
	}

	elections, err := s.repository.GetElections(validateLimit, validateOffset, user.ID)
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
