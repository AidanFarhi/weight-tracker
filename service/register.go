package service

import (
	"errors"
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
	// check if user already exists
	err := as.ur.CheckIfUserExists(email)
	if err == nil {
		return "", errors.New("user already exists")
	}

	// create a user
	userId, err := as.ur.CreateUser(email, password)
	if err != nil {
		return "", errors.New("error creating user")
	}

	// create a session with the user id using session repo
	sessionId, err := as.sr.CreateSession(userId)
	if err != nil {
		return "", err
	}

	return sessionId, nil
}
