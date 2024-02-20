package storage

import "errors"

var (
	ErrUserExist = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
	ErrAppNotFound = errors.New("app not found")
)