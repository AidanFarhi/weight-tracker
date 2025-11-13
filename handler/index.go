package handler

import "net/http"

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (ih *IndexHandler) GetIndex(w http.ResponseWriter, r *http.Request) {
	// implement redirect in middlware layer
	http.Redirect(w, r, "/login", http.StatusFound)
}
