package routers

import (
	"github.com/gorilla/mux"
	"github.com/ifeanyilawrence/go-task-api/controllers"
	"github.com/ifeanyilawrence/go-task-api/middleware"
)

//SetTaskRoutes : Sets the task routes
func SetTaskRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/tasks", middleware.AuthMiddleware(controllers.GetTasks)).Methods("GET")
	router.HandleFunc("/api/tasks", middleware.AuthMiddleware(controllers.CreateTask)).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", middleware.AuthMiddleware(controllers.GetSingleTask)).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", middleware.AuthMiddleware(controllers.UpdateTask)).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", middleware.AuthMiddleware(controllers.DeleteTask)).Methods("DELETE")

	return router
}
