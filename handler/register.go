package handler

import (
	"net/http"
	"text/template"
)

type RegisterHandler struct{}

func NewRegisterHandler() *RegisterHandler {
	return &RegisterHandler{}
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
