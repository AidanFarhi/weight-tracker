package handler

import (
	"net/http"
	"text/template"
	"weight-tracker/model"
	"weight-tracker/service"
)

const (
	emptyFieldErrorMessage          = "Please provide an email and a valid password."
	passwordNotMatchingErrorMessage = "Passwords do not match."
	userAlreadyExistsErrorMessage   = "A user with that email address already exists."
)

type RegisterHandler struct {
	rs *service.RegisterService
}

func NewRegisterHandler(rs *service.RegisterService) *RegisterHandler {
	return &RegisterHandler{rs}
}

func (rh *RegisterHandler) GetRegister(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"./web/templates/base.html",
		"./web/templates/pages/register.html",
	)
	if err != nil {
		http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "register", nil)
}

func (rh *RegisterHandler) PostRegister(w http.ResponseWriter, r *http.Request) {
	// get the credentials from the form
	email := r.FormValue("register-email")
	password := r.FormValue("register-password")
	passwordRepeat := r.FormValue("register-password-repeat")

	// check if credentials are missing
	if email == "" || password == "" || passwordRepeat == "" {
		t, err := template.ParseFiles(
			"./web/templates/base.html",
			"./web/templates/pages/register.html",
		)
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "register", model.RegisterPageData{
			RegisterError:        true,
			RegisterErrorMessage: emptyFieldErrorMessage,
		})
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// check if passwords match
	if password != passwordRepeat {
		t, err := template.ParseFiles(
			"./web/templates/base.html",
			"./web/templates/pages/register.html",
		)
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "register", model.RegisterPageData{
			RegisterError:        true,
			RegisterErrorMessage: passwordNotMatchingErrorMessage,
		})
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// attempt to register a user
	_, err := rh.rs.Register(email, password)

	if err != nil {
		if err.Error() == "user already exists" {
			t, err := template.ParseFiles(
				"./web/templates/base.html",
				"./web/templates/pages/register.html",
			)
			if err != nil {
				http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
				return
			}
			err = t.ExecuteTemplate(w, "register", model.RegisterPageData{
				RegisterError:        true,
				RegisterErrorMessage: userAlreadyExistsErrorMessage,
			})
			if err != nil {
				http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
			}
			return
		}
		http.Error(w, "Error while registering a user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// generate a session using the userId
	sessionId := "sessionId1234"

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
