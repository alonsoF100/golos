package postgres

import (
	"time"

	"github.com/alonsoF100/golos/internal/models"
)

func (r Repository) CreateVote(uuid, userID, voteVariantID string, createdAt time.Time, updatedAt time.Time) (*models.Vote, error) {
	return nil, nil
}

func (r Repository) GetVote(uuid string) (*models.Vote, error) {
	return nil, nil
}

func (r Repository) GetUserVotes(userID string) ([]*models.Vote, error) {
	return nil, nil
}

func (r Repository) GetVariantVotes(voteVariantID string) ([]*models.Vote, error) {
	return nil, nil
}

func (r Repository) DeleteVote(uuid string) error {
	return nil
}

func (r Repository) PathVote(uuid string, userID, voteVariantID *string, updatedAt time.Time) (*models.Vote, error) {
	return nil, nil
}
