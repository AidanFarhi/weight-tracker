package repo

import "errors"

var (
	ErrSQLDBIssue         = errors.New("error while trying to query DB")
	ErrUserNotFound       = errors.New("could not find user")
	ErrDuplicateSessionId = errors.New("session id already exists")
	ErrNoRecordsFound     = errors.New("no records found")
)
