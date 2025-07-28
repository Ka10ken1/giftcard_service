package controller

import (
	"net/http"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/health", healthCheck).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()

	RegisterUserRoutes(api)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Healthy"))
}
