package main

import (
	"log"
	"net/http"
	"weight-tracker/handler"
	"weight-tracker/middleware"
	"weight-tracker/repo"
	"weight-tracker/service"
)

func main() {

	// init repos
	ur := repo.NewUserRepo()
	sr := repo.NewSessionRepo()
	wr := repo.NewWeightRepo()

	// init services
	as := service.NewAuthService(ur, sr)
	rs := service.NewRegisterService(ur, sr)
	ws := service.NewWeightService(wr)

	// init middleware
	am := middleware.NewAuthMiddleware(as)

	// init handlers
	hh := handler.NewHomeHandler()
	ah := handler.NewAuthHandler(as)
	rh := handler.NewRegisterHandler(rs)
	wh := handler.NewWeightHandler(ws)

	// create multiplexer
	mux := http.NewServeMux()

	// register routes with multiplexer
	mux.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	mux.HandleFunc("/", am.RequireAuth(hh.GetHome))
	mux.HandleFunc("GET /login", am.RedirectIfLoggedIn(ah.GetLogin))
	mux.HandleFunc("POST /login", am.RedirectIfLoggedIn(ah.PostLogin))
	mux.HandleFunc("GET /register", am.RedirectIfLoggedIn(rh.GetRegister))
	mux.HandleFunc("POST /register", am.RedirectIfLoggedIn(rh.PostRegister))
	mux.HandleFunc("GET /api/daily-weights", am.RequireAuth(wh.GetDailyWeights))

	// create server
	server := http.Server{
		Addr:    ":8899",
		Handler: mux,
	}

	// start server
	log.Fatal(server.ListenAndServe())
}
