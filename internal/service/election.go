package service

import (
	"time"

	apperrors "github.com/alonsoF100/golos/internal/erorrs"
	"github.com/alonsoF100/golos/internal/models"
	"github.com/google/uuid"
)

func (s ElectionService) CreateElection(userID string, name string, description string) (*models.Election, error) {
	now := time.Now()
	id := uuid.New().String()

	election, err := s.electionRepository.CreateElection(id, userID, name, description, now, now)
	if err != nil {
		return nil, err
	}

	return election, nil
}

func (s ElectionService) GetElection(uuid string) (*models.Election, error) {
	election, err := s.electionRepository.GetElection(uuid)
	if err != nil {
		return nil, err
	}

	return election, nil
}

func (s ElectionService) DeleteElection(uuid string) error {
	err := s.electionRepository.DeleteElection(uuid)
	if err != nil {
		return err
	}

	return nil
}

func (s ElectionService) PatchElection(uuid string, userID, name, description *string) (*models.Election, error) {
	now := time.Now()
	if userID == nil && name == nil && description == nil {
		return nil, apperrors.ErrNothingToChange
	}

	election, err := s.electionRepository.PatchElection(uuid, userID, name, description, now)
	if err != nil {
		return nil, err
	}

	return election, nil
}
