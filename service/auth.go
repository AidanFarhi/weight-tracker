package service

import "weight-tracker/repo"

type AuthService struct {
	ur *repo.UserRepo
}

func NewAuthService(ur *repo.UserRepo) *AuthService {
	return &AuthService{ur}
}

func (as *AuthService) Login(email, password string) (string, error) {
	// check user credentials
	_, err := as.ur.GetIdForUser(email, password)
	if err != nil {
		return "", err
	}
	// attempt to create a session using userId and session repo

	// return newly created session id
	return "sessionid123456", nil
}

func (as *AuthService) Logout() {}
