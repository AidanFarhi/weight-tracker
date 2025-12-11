package service

import (
	"weight-tracker/repo"
)

type AuthService struct {
	ur UserRepo
	sr SessionRepo
}

func NewAuthService(ur UserRepo, sr SessionRepo) *AuthService {
	return &AuthService{
		ur: ur,
		sr: sr,
	}
}

func (as *AuthService) Login(email, password string) (string, error) {
	user, err := as.ur.GetUserByEmail(email)
	if err == repo.ErrNoRecordsFound {
		return "", ErrEmailNotFound
	}
	if err != nil {
		return "", ErrDBIssue
	}
	passwordValid := VerifyPassword(password, user.Password)
	if !passwordValid {
		return "", ErrIncorrectPassword
	}
	// create a session
	sessionId := GenerateSessionId()
	err = as.sr.CreateSession(sessionId, user.Id)
	// try again if there is a duplicate session id (almost impossible)
	if err == repo.ErrDuplicateSessionId {
		sessionId = GenerateSessionId()
		err = as.sr.CreateSession(sessionId, user.Id)
	}
	if err != nil {
		return "", ErrDBIssue
	}
	// return newly created session id
	return sessionId, nil
}

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
