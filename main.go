package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"weight-tracker/handler"
)

func HandleError(operation string, err error) {
	if err != nil {
		fmt.Println("error "+operation+":", err.Error())
		os.Exit(1)
	}
}

func main() {

	// init handlers
	ih := handler.NewIndexHandler()
	lh := handler.NewLoginHandler()
	rh := handler.NewRegisterHandler()

	// create multiplexer
	mux := http.NewServeMux()

	// register routes with multiplexer
	mux.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	mux.HandleFunc("/", ih.GetIndex)
	mux.HandleFunc("GET /login", lh.GetLogin)
	mux.HandleFunc("GET /register", rh.GetRegister)

	// create server
	server := http.Server{
		Addr:    ":8899",
		Handler: mux,
	}

	// start server
	log.Fatal(server.ListenAndServe())
}
