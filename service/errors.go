package service

import "errors"

var (
	ErrDBIssue           = errors.New("error while performing DB operation")
	ErrNoSessionFound    = errors.New("no session found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrDateParseIssue    = errors.New("error parsing date")
	ErrEmailNotFound     = errors.New("email not found")
	ErrIncorrectPassword = errors.New("password does not match")
	ErrUserNotFound      = errors.New("user not found")
	ErrHashingIssue      = errors.New("error while trying to hash")
)
