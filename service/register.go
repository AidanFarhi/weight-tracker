package service

import (
	"errors"
	"weight-tracker/repo"
)

type RegisterService struct {
	ur *repo.UserRepo
}

func NewRegisterService(ur *repo.UserRepo) *RegisterService {
	return &RegisterService{ur}
}

func (as *RegisterService) Register(email, password string) (string, error) {
	// check if user already exists
	err := as.ur.CheckIfUserExists(email)
	if err == nil {
		return "", errors.New("user already exists")
	}

	// create a user
	_, err = as.ur.CreateUser(email, password)
	if err != nil {
		return "", errors.New("error creating user")
	}

	// create a session with the user id using session repo
	sessionId := "sessionId1234"

	return sessionId, nil
}
