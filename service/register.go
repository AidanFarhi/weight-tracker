package service

import (
	"weight-tracker/repo"
)

type RegisterService struct {
	ur UserRepo
	sr SessionRepo
}

func NewRegisterService(ur UserRepo, sr SessionRepo) *RegisterService {
	return &RegisterService{
		ur: ur,
		sr: sr,
	}
}

func (as *RegisterService) Register(email, password string) (string, error) {
	userExists, err := as.ur.CheckIfUserExists(email)
	if err != nil {
		return "", ErrDBIssue
	}
	if userExists {
		return "", ErrUserAlreadyExists
	}
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return "", ErrHashingIssue
	}
	userId, err := as.ur.CreateUser(email, hashedPassword)
	if err != nil {
		return "", ErrDBIssue
	}
	sessionId := GenerateSessionId()
	err = as.sr.CreateSession(sessionId, userId)
	if err == repo.ErrDuplicateSessionId {
		sessionId = GenerateSessionId()
		err = as.sr.CreateSession(sessionId, userId)
	}
	if err != nil {
		return "", ErrDBIssue
	}
	return sessionId, nil
}
