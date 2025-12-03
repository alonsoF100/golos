package service

import (
	"time"

	"github.com/alonsoF100/golos/internal/models"
	"github.com/google/uuid"
)

func (s Service) CreateVoteVariant(electionID, name string) (*models.VoteVariant, error) {
	now := time.Now()
	id := uuid.New().String()

	voteVariant, err := s.repository.CreateVoteVariant(id, electionID, name, now, now)
	if err != nil {
		return nil, err
	}

	return voteVariant, nil
}

// //cтранная функция вроде нелогичная, не забыть обсудить
// добавить квэрипараметры
func (s Service) GetVoteVariants(electionID string) ([]*models.VoteVariant, error) {
	voteVariants, err := s.repository.GetVoteVariants(electionID)
	if err != nil {
		return nil, err
	}

	return voteVariants, nil
}

func (s Service) GetVoteVariant(uuid string) (*models.VoteVariant, error) {
	voteVariant, err := s.repository.GetVoteVariant(uuid)
	if err != nil {
		return nil, err
	}

	return voteVariant, nil
}

func (s Service) DeleteVoteVariant(uuid string) error {
	err := s.repository.DeleteVoteVariant(uuid)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) UpdateVoteVariant(uuid string, name string) (*models.VoteVariant, error) {
	now := time.Now()

	voteVariant, err := s.repository.UpdateVoteVariant(uuid, name, now)
	if err != nil {
		return nil, err
	}

	return voteVariant, nil
}
