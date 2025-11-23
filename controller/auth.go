package controller

import (
	"net/http"
	"weight-tracker/model"
	"weight-tracker/service"
)

type AuthController struct {
	as *service.AuthService
}

func NewAuthController(as *service.AuthService) *AuthController {
	return &AuthController{as}
}

func (ac *AuthController) GetLogin(w http.ResponseWriter, r *http.Request) {
	err := RenderPage(w, "login", model.SimplePageData{
		HasError:     false,
		ErrorMessage: "",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ac *AuthController) PostLogin(w http.ResponseWriter, r *http.Request) {
	// get the credentials from the form
	email := r.FormValue("email")
	password := r.FormValue("password")

	// check if credentials are missing
	if email == "" || password == "" {
		err := RenderPage(w, "login", model.SimplePageData{
			HasError:     true,
			ErrorMessage: EmptyEmailOrPasswordMessage,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// attempt a login with login service
	sessionId, err := ac.as.Login(email, password)

	// check if the error is due to no records found
	if err != nil {
		err := RenderPage(w, "login", model.SimplePageData{
			HasError:     true,
			ErrorMessage: InvalidCredentialsMessage,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// store session id in cookie
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		// Expires:  time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(w, cookie)

	// redirect to home page "/"
	http.Redirect(w, r, "/", http.StatusFound)
}
