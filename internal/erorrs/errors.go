package apperrors

import "errors"

var (
	ErrUserAlreadyExist = errors.New("user aready exist")
)