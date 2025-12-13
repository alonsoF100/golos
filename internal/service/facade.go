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
