package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)


func RegisterUserRoutes(r *mux.Router) {
	user := r.PathPrefix("/users").Subrouter()

	user.HandleFunc("", getAllUsers).Methods("GET")

	user.HandleFunc("/{id}", getUserByID).Methods("GET")
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of users"))
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Write([]byte("User ID: " + id))
}
