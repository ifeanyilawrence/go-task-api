package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ifeanyilawrence/go-task-api/controllers"
)

func main() {
	hndlr := mux.NewRouter()

	hndlr.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	hndlr.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", hndlr))
}
