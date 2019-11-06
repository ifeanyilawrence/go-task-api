package routers

import (
	"github.com/gorilla/mux"
	"github.com/ifeanyilawrence/go-task-api/controllers"
)

//SetAuthRoutes : Sets the auth routes
func SetAuthRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/auth/getToken", controllers.GetToken).Methods("POST")
	router.HandleFunc("/api/auth/profile", controllers.Profile).Methods("POST")

	return router
}
