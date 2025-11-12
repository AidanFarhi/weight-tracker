package handler

import "net/http"

func GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<p>Hello World</p>"))
}
