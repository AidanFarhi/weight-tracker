package middleware

import (
	"context"
	"net/http"
	"weight-tracker/service"
)

type AuthMiddleware struct {
	as *service.AuthService
}

func NewAuthMiddleware(as *service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{as}
}

func (am *AuthMiddleware) RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		userId, err := am.as.ValidateSession(cookie.Value)
		if err == service.ErrNoSessionFound {
			// clear the cookie and redirect to login
			expiredCookie := &http.Cookie{
				Name:     "session_id",
				Value:    "",
				Path:     "/",
				MaxAge:   -1,
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteStrictMode,
			}
			http.SetCookie(w, expiredCookie)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		if err == service.ErrDBIssue {
			http.Error(w, "Error validating session: "+err.Error(), http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), UserIDKey, userId)
		next(w, r.WithContext(ctx))
	}
}

func (am *AuthMiddleware) RedirectIfLoggedIn(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			next(w, r)
			return
		}
		_, err = am.as.ValidateSession(cookie.Value)
		if err == service.ErrNoSessionFound {
			// clear the cookie and redirect to login
			expiredCookie := &http.Cookie{
				Name:     "session_id",
				Value:    "",
				Path:     "/",
				MaxAge:   -1,
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteStrictMode,
			}
			http.SetCookie(w, expiredCookie)
			next(w, r)
			return
		}
		if err == service.ErrDBIssue {
			http.Error(w, "Error validating session: "+err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
