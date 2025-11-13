package handler

import (
	"net/http"
	"text/template"
	"weight-tracker/model"
	"weight-tracker/service"
)

const (
	invalidCredentialsMessage   = "Sorry. We could not find a match for your email and/or password."
	emptyEmailOrPasswordMessage = "Please provide both email and password."
)

type AuthHandler struct {
	as *service.AuthService
}

func NewAuthHandler(as *service.AuthService) *AuthHandler {
	return &AuthHandler{as}
}

func (ah *AuthHandler) GetLogin(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"./web/templates/base.html",
		"./web/templates/pages/login.html",
	)
	if err != nil {
		http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "login", nil)
	if err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ah *AuthHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
	// get the credentials from the form
	email := r.FormValue("email")
	password := r.FormValue("password")

	// check if credentials are missing
	if email == "" || password == "" {
		t, err := template.ParseFiles(
			"./web/templates/base.html",
			"./web/templates/pages/login.html",
		)
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "login", model.LoginPageData{
			LoginError:        true,
			LoginErrorMessage: emptyEmailOrPasswordMessage,
		})
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// attempt a login with login service
	sessionId, err := ah.as.Login(email, password)

	// check if the error is due to no records found
	if err != nil {
		t, err := template.ParseFiles(
			"./web/templates/base.html",
			"./web/templates/pages/login.html",
		)
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "login", model.LoginPageData{
			LoginError:        true,
			LoginErrorMessage: invalidCredentialsMessage,
		})
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
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
