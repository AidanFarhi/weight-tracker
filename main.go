package main

import (
	"context"
	"log"
	"net/http"
	"weight-tracker/config"
	"weight-tracker/controller"
	"weight-tracker/middleware"
	"weight-tracker/repo"
	"weight-tracker/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	// load config
	c, _ := config.LoadConfig()

	// create db connection pool
	pool, err := pgxpool.New(context.Background(), c.DBURI)
	if err != nil {
		log.Fatal("error creating DB connection pool")
	}
	defer pool.Close()

	// init repos
	ur := repo.NewUserRepo(pool)
	sr := repo.NewSessionRepo(pool)
	wr := repo.NewWeightRepo(pool)

	// init services
	as := service.NewAuthService(ur, sr)
	rs := service.NewRegisterService(ur, sr)
	ws := service.NewWeightService(wr)

	// init middleware
	am := middleware.NewAuthMiddleware(as)

	// init handlers
	hc := controller.NewHomeController(ws)
	ac := controller.NewAuthController(as)
	rc := controller.NewRegisterController(rs)
	wc := controller.NewWeightController(ws)

	// create multiplexer
	mux := http.NewServeMux()

	// register routes with multiplexer
	mux.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	mux.HandleFunc("/", am.RequireAuth(hc.GetHome))
	mux.HandleFunc("GET /login", am.RedirectIfLoggedIn(ac.GetLogin))
	mux.HandleFunc("POST /login", am.RedirectIfLoggedIn(ac.PostLogin))
	mux.HandleFunc("GET /register", am.RedirectIfLoggedIn(rc.GetRegister))
	mux.HandleFunc("POST /register", am.RedirectIfLoggedIn(rc.PostRegister))
	mux.HandleFunc("GET /daily-weight-entry", am.RequireAuth(wc.GetDailyWeightEntry))
	mux.HandleFunc("POST /daily-weight-entry", am.RequireAuth(wc.PostDailyWeightEntry))
	mux.HandleFunc("GET /api/daily-weights", am.RequireAuth(wc.GetDailyWeights))
	mux.HandleFunc("GET /target-weight-entry", am.RequireAuth(wc.GetTargetWeightEntry))
	mux.HandleFunc("POST /target-weight-entry", am.RequireAuth(wc.PostTargetWeightEntry))

	// create server
	server := http.Server{
		Addr:    ":8899",
		Handler: mux,
	}

	// start server
	log.Fatal(server.ListenAndServe())
}
