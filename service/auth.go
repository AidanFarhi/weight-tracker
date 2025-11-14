package service

import "weight-tracker/repo"

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

	// attempt to create a session using userId and session repo
	sessionId, err := as.sr.CreateSession(userId)
	if err != nil {
		return "", err
	}

	// return newly created session id
	return sessionId, nil
}

func (as *AuthService) Logout() {}

func (as *AuthService) ValidateSession(sessionId string) (int, error) {
	userId, err := as.sr.GetUserIdForSession(sessionId)
	if err != nil {
		return -1, err
	}
	if userId == -1 {
		return -1, ErrNoSessionFound
	}
	return userId, nil
}
