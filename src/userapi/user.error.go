package userapi

import "errors"

var (
	ErrNotImplemented = errors.New("not yet implemented")
	ErrUserDoesNotExist = errors.New("user does not exist")
	ErrUserAlreadyExists = errors.New("user already exists")
)