package handler

import "net/http"

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (ih *HomeHandler) GetHome(w http.ResponseWriter, r *http.Request) {
	// implement redirect in middlware layer
	http.Redirect(w, r, "/login", http.StatusFound)
}
