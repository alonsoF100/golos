package service

import (
	"time"

	"github.com/alonsoF100/golos/internal/models"
	"github.com/google/uuid"
)

func (s VoteVariantService) CreateVoteVariant(electionID, name string) (*models.VoteVariant, error) {
	now := time.Now()
	id := uuid.New().String()

	voteVariant, err := s.voteVariantRepository.CreateVoteVariant(id, electionID, name, now, now)
	if err != nil {
		return nil, err
	}

	return voteVariant, nil
}

func (s VoteVariantService) GetVoteVariants(electionID string) ([]*models.VoteVariant, error) {
	voteVariants, err := s.voteVariantRepository.GetVoteVariants(electionID)
	if err != nil {
		return nil, err
	}

	return voteVariants, nil
}

func (s VoteVariantService) GetVoteVariant(uuid string) (*models.VoteVariant, error) {
	voteVariant, err := s.voteVariantRepository.GetVoteVariant(uuid)
	if err != nil {
		return nil, err
	}

	return voteVariant, nil
}

func (s VoteVariantService) DeleteVoteVariant(uuid string) error {
	err := s.voteVariantRepository.DeleteVoteVariant(uuid)
	if err != nil {
		return err
	}

	return nil
}

func (s VoteVariantService) UpdateVoteVariant(uuid string, name string) (*models.VoteVariant, error) {
	now := time.Now()

	voteVariant, err := s.voteVariantRepository.UpdateVoteVariant(uuid, name, now)
	if err != nil {
		return nil, err
	}

	return voteVariant, nil
}
