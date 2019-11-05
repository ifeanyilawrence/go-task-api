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
	hndlr.HandleFunc("/api/users/{id}", controllers.GetSingleUser).Methods("GET")
	hndlr.HandleFunc("/api/users/{id}", controllers.UpdateUser).Methods("PUT")
	hndlr.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", hndlr))
}
