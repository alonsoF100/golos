package service

import (
	"time"

	apperrors "github.com/alonsoF100/golos/internal/erorrs"
	"github.com/alonsoF100/golos/internal/models"
	"github.com/google/uuid"
)

func (s VoteService) CreateVote(userID, voteVariantID string) (*models.Vote, error) {
	id := uuid.New().String()
	now := time.Now()

	vote, err := s.voteRepository.CreateVote(id, userID, voteVariantID, now, now)
	if err != nil {
		return nil, err
	}

	return vote, nil
}

func (s VoteService) GetVote(voteID string) (*models.Vote, error) {
	vote, err := s.voteRepository.GetVote(voteID)
	if err != nil {
		return nil, err
	}

	return vote, nil
}

func (s VoteService) DeleteVote(voteID string) error {
	err := s.voteRepository.DeleteVote(voteID)
	if err != nil {
		return err
	}

	return nil
}

func (s VoteService) PatchVote(voteID string, userID, voteVariantID *string) (*models.Vote, error) {
	now := time.Now()

	if userID == nil && voteVariantID == nil {
		return nil, apperrors.ErrNothingToChange
	}

	vote, err := s.voteRepository.PatchVote(voteID, userID, voteVariantID, now)
	if err != nil {
		return nil, err
	}

	return vote, nil
}
