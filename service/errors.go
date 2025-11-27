package service

import "errors"

var (
	ErrDBIssue           = errors.New("error while performing DB operation")
	ErrNoSessionFound    = errors.New("no session found")
	ErrUserAlreadyExists = errors.New("user already exists")
)
