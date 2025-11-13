package handler

import (
	"net/http"
	"text/template"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (ih *HomeHandler) GetHome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"./web/templates/base.html",
		"./web/templates/pages/home.html",
	)
	if err != nil {
		http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "home", nil)
	if err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
	}
}
