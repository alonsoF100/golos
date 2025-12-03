package service

import (
	"time"

	apperrors "github.com/alonsoF100/golos/internal/erorrs"
	"github.com/alonsoF100/golos/internal/models"
	"github.com/google/uuid"
)

func (s Service) CreateVote(userID, voteVariantID string) (*models.Vote, error) {
	id := uuid.New().String()
	now := time.Now()

	vote, err := s.repository.CreateVote(id, userID, voteVariantID, now, now)
	if err != nil {
		return nil, err
	}

	return vote, nil
}

func (s Service) GetVote(voteID string) (*models.Vote, error) {
	vote, err := s.repository.GetVote(voteID)
	if err != nil {
		return nil, err
	}

	return vote, nil
}

func (s Service) GetUserVotes(userID string) (*[]models.Vote, error) {
	return nil, nil
}

func (s Service) GetVariantVotes(voteVariantID string) (*[]models.Vote, error) {
	return nil, nil
}

func (s Service) DeleteVote(voteID string) error {
	err := s.repository.DeleteVote(voteID)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) PatchVote(voteID string, userID, voteVariantID *string) (*models.Vote, error) {
	now := time.Now()

	if userID == nil && voteVariantID == nil {
		return nil, apperrors.ErrNothingToChange
	}

	vote, err := s.repository.PatchVote(voteID, userID, voteVariantID, now)
	if err != nil {
		return nil, err
	}

	return vote, nil
}
