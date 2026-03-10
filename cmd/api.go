package main

import (
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
)

type application struct {
	addr         string
	dsn          string
	dbConnection *pgx.Conn
}

func (app *application) registerControllers() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Word!\n"))
	})

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All good\n"))
	})

	return mux
}

func (app *application) getServer(handler http.Handler) *http.Server {
	server := http.Server{
		Addr:         app.addr,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second * 10,
		Handler:      handler,
	}

	return &server
}
