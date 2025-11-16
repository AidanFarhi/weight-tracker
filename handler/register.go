package handler

import (
	"net/http"
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
	err := RenderPage(w, "register", model.SimplePageData{
		HasError:     false,
		ErrorMessage: "",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rh *RegisterHandler) PostRegister(w http.ResponseWriter, r *http.Request) {
	// get the credentials from the form
	email := r.FormValue("register-email")
	password := r.FormValue("register-password")
	passwordRepeat := r.FormValue("register-password-repeat")

	// check if credentials are missing
	if email == "" || password == "" || passwordRepeat == "" {
		err := RenderPage(w, "register", model.SimplePageData{
			HasError:     true,
			ErrorMessage: emptyFieldErrorMessage,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// check if passwords match
	if password != passwordRepeat {
		err := RenderPage(w, "register", model.SimplePageData{
			HasError:     true,
			ErrorMessage: passwordNotMatchingErrorMessage,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// attempt to register a user
	sessionId, err := rh.rs.Register(email, password)

	if err.Error() == "user already exists" {
		err := RenderPage(w, "register", model.SimplePageData{
			HasError:     true,
			ErrorMessage: userAlreadyExistsErrorMessage,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// some other error occured
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
