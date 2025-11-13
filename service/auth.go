package service

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (as *AuthService) Login() {}

func (as *AuthService) Logout() {}
