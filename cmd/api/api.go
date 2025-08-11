package main

import (
	"log"
	"net/http"
	"time"
)
type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(http.MethodGet + "/v1/health", app.health()) 
   return mux
}

func (app *application) run(mux *http.ServeMux) error {
	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Println("Starting server on", app.config.addr)
	return srv.ListenAndServe()

}