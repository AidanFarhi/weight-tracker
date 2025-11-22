package repo

import "errors"

var (
	ErrSQLDBIssue   = errors.New("error while trying to query DB")
	ErrUserNotFound = errors.New("could not find user")
)
