package main

import (
	"log"
	"net/http"

	"github.com/Ka10ken1/giftcard_service/internal/config"
	"github.com/Ka10ken1/giftcard_service/internal/controller"
	"github.com/gorilla/mux"
)

func main(){

	config.ConnectDatabase()

	r := mux.NewRouter()

	controller.RegisterRoutes(r)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
