package main

import (
	"github/onunkwor/social/internal/store"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// application holds the configuration and dependencies for the API.
type application struct {
	config config
	store  store.Storage
}

// config holds the configuration for the application.
type config struct {
	addr string
	db   dbConfig
}

// dbConfig holds the database configuration.
type dbConfig struct {
	dsn string
}

// healthCheckHandler is a simple handler to check if the API is running.
func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

// healthCheckHandler responds with a simple message indicating the API is healthy.
func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Println("Starting server on", app.config.addr)
	return srv.ListenAndServe()

}
