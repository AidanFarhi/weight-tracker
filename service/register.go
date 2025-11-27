package service

import (
	"weight-tracker/repo"
)

type RegisterService struct {
	ur *repo.UserRepo
	sr *repo.SessionRepo
}

func NewRegisterService(ur *repo.UserRepo, sr *repo.SessionRepo) *RegisterService {
	return &RegisterService{
		ur: ur,
		sr: sr,
	}
}

func (as *RegisterService) Register(email, password string) (string, error) {
	userExists, err := as.ur.CheckIfUserExists(email)

	if err != nil {
		return "", err
	}

	if userExists {
		return "", ErrUserAlreadyExists
	}

	// create a user
	// TODO: hash the password using secure crypto stuff
	userId, err := as.ur.CreateUser(email, password)
	if err != nil {
		return "", err
	}

	// create a session id
	sessionId := GenerateSessionId()

	// create a session with the user id using session repo
	err = as.sr.CreateSession(sessionId, userId)

	// try again if there is a duplicate session id (almost impossible)
	if err == repo.ErrDuplicateSessionId {
		sessionId = GenerateSessionId()
		err = as.sr.CreateSession(sessionId, userId)
	}

	if err != nil {
		return "", ErrDBIssue
	}

	return sessionId, nil
}
