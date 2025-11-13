package main

import (
	"log"
	"net/http"
	"weight-tracker/handler"
	"weight-tracker/repo"
	"weight-tracker/service"
)

func main() {

	// init repos
	ur := repo.NewUserRepo()

	// init services
	us := service.NewAuthService(ur)

	// init handlers
	hh := handler.NewHomeHandler()
	ah := handler.NewAuthHandler(us)
	rh := handler.NewRegisterHandler()

	// create multiplexer
	mux := http.NewServeMux()

	// register routes with multiplexer
	mux.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	mux.HandleFunc("/", hh.GetHome)
	mux.HandleFunc("GET /login", ah.GetLogin)
	mux.HandleFunc("POST /login", ah.PostLogin)
	mux.HandleFunc("GET /register", rh.GetRegister)

	// create server
	server := http.Server{
		Addr:    ":8899",
		Handler: mux,
	}

	// start server
	log.Fatal(server.ListenAndServe())
}
