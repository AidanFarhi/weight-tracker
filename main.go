package main

import (
	"log"
	"net/http"
	"weight-tracker/handler"
)

func main() {

	fs := http.StripPrefix("/web/", http.FileServer(http.Dir("./web")))

	mux := http.NewServeMux()

	mux.Handle("/web/", fs)
	mux.HandleFunc("/", handler.GetIndex)

	server := http.Server{
		Addr:    ":8899",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
