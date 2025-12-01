package apperrors

import "errors"

var (
	// User errors
	ErrUserAlreadyExist     = errors.New("user already exist")
	ErrUserNotFound         = errors.New("user not found")
	ErrNothingToChange      = errors.New("nothing to change")
	ErrFailedToHashPassword = errors.New("failed to hash password")

	// Election errors
	ErrElectionAlreadyExist = errors.New("election already exist")
	ErrElectionNotFound     = errors.New("election not found")

	// Vote Variant errors
	ErrVoteVariantAlreadyExist = errors.New("vote variant already exist")
	ErrVoteVariantNotFound     = errors.New("vote variant not found")
)
