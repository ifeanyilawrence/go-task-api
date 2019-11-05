package routers

import (
	"github.com/gorilla/mux"
	"github.com/ifeanyilawrence/go-task-api/controllers"
)

//SetTaskRoutes : Sets the task routes
func SetTaskRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", controllers.GetSingleTask).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", controllers.DeleteTask).Methods("DELETE")

	return router
}
