package service

import "github.com/alonsoF100/golos/internal/models"

func (s Service) CreateVote(userID, voteVariantID string) (*models.Vote, error) {
	return nil, nil
}

func (s Service) GetVote(voteID string) (*models.Vote, error) {
	return nil, nil
}

func (s Service) GetUserVotes(userID string) (*[]models.Vote, error) {
	return nil, nil
}

func (s Service) GetVariantVotes(voteVariantID string) (*[]models.Vote, error) {
	return nil, nil
}

func (s Service) DeleteVote(voteID string) error {
	return nil
}

func (s Service) PatchVote(voteID string, userID, voteVariantID *string) (*models.Vote, error) {
	return nil, nil
}
