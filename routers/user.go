package routers

import (
	"github.com/gorilla/mux"
	"github.com/ifeanyilawrence/go-task-api/controllers"
)

//SetUserRoutes : sets the user routes
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", controllers.GetSingleUser).Methods("GET")
	router.HandleFunc("/api/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods("DELETE")

	return router
}
