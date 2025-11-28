package apperrors

import "errors"

var (
	ErrUserAlreadyExist = errors.New("user aready exist")
	ErrUserNotFound     = errors.New("user not found")
	ErrNothingToChange = errors.New("nothing to change")
)
