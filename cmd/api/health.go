package main

import "net/http"

func (app *application) health () http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is Alive"))
	}
}