package handler

import (
	"net/http"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (ih *HomeHandler) GetHome(w http.ResponseWriter, r *http.Request) {
	err := RenderPage(w, "home", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
