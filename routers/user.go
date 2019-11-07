package routers

import (
	"github.com/gorilla/mux"
	"github.com/ifeanyilawrence/go-task-api/controllers"
	"github.com/ifeanyilawrence/go-task-api/middleware"
)

//SetUserRoutes : sets the user routes
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/users", middleware.AuthMiddleware(controllers.GetUsers)).Methods("GET")
	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", middleware.AuthMiddleware(controllers.GetSingleUser)).Methods("GET")
	router.HandleFunc("/api/users/{id}", middleware.AuthMiddleware(controllers.UpdateUser)).Methods("PUT")
	router.HandleFunc("/api/users/{id}", middleware.AuthMiddleware(controllers.DeleteUser)).Methods("DELETE")

	return router
}
