package handler

import (
	"net/http"
	"text/template"
)

type LoginHandler struct{}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{}
}

func (lh *LoginHandler) GetLogin(w http.ResponseWriter, r *http.Request) {
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
