package service

import (
	"weight-tracker/repo"
)

type AuthService struct {
	ur *repo.UserRepo
	sr *repo.SessionRepo
}

func NewAuthService(ur *repo.UserRepo, sr *repo.SessionRepo) *AuthService {
	return &AuthService{
		ur: ur,
		sr: sr,
	}
}

func (as *AuthService) Login(email, password string) (string, error) {
	// check user credentials
	userId, err := as.ur.GetIdForUser(email, password)
	if err != nil {
		return "", err
	}

	// create a session
	sessionId := GenerateSessionId()
	err = as.sr.CreateSession(sessionId, userId)

	// try again if there is a duplicate session id (almost impossible)
	if err == repo.ErrDuplicateSessionId {
		sessionId = GenerateSessionId()
		err = as.sr.CreateSession(sessionId, userId)
	}

	if err != nil {
		return "", ErrDBIssue
	}

	// return newly created session id
	return sessionId, nil
}

func (as *AuthService) Logout() {}

func (as *AuthService) ValidateSession(sessionId string) (int, error) {
	userId, err := as.sr.GetUserIdForSession(sessionId)
	if err == repo.ErrNoRecordsFound {
		return 0, ErrNoSessionFound
	}
	if err != nil {
		return 0, ErrDBIssue
	}
	return userId, nil
}
