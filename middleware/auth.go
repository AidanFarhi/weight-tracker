package middleware

import "weight-tracker/service"

type AuthMiddleware struct {
	as *service.AuthService
}

func NewAuthMiddleware(as *service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{as}
}
