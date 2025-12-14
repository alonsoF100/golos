package service

import (
	apperrors "github.com/alonsoF100/golos/internal/erorrs"
	"github.com/alonsoF100/golos/internal/models"
)

func (s Service) GetElections(limit, offset int, nickname string) ([]*models.Election, error) {
	validateLimit := validateLimit(limit)
	validateOffset := validateOffset(offset)

	user, err := s.UserService.userRepository.GetUserByNickname(nickname)
	if err != nil {
		return nil, apperrors.ErrUserNotFound
	}

	elections, err := s.ElectionService.electionRepository.GetElections(validateLimit, validateOffset, user.ID)
	if err != nil {
		return nil, err
	}

	return elections, nil
}

func (s Service) GetUserVotes(nickname, electionID string, limit int, offset int) ([]*models.Vote, error) {
	validLimit := validateLimit(limit)
	validOffset := validateOffset(offset)

	user, err := s.UserService.userRepository.GetUserByNickname(nickname)
	if err != nil {
		return nil, apperrors.ErrUserNotFound
	}

	voteVariants, err := s.VoteVariantService.voteVariantRepository.GetVoteVariants(electionID)
	if err != nil {
		return nil, err
	}

	voteVariantIDs := make([]string, 0, len(voteVariants))
	for _, voteVariant := range voteVariants {
		voteVariantIDs = append(voteVariantIDs, voteVariant.ID)
	}

	votes, err := s.VoteService.voteRepository.GetUserVotes(user.ID, voteVariantIDs, validLimit, validOffset)
	if err != nil {
		return nil, err
	}

	return votes, nil
}

func (s Service) GetVariantVotes(voteVariantID string) ([]*models.Vote, error) {
	return nil, nil
}
