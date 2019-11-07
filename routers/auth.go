package routers

import (
	"github.com/gorilla/mux"
	"github.com/ifeanyilawrence/go-task-api/controllers"
	"github.com/ifeanyilawrence/go-task-api/middleware"
)

//SetAuthRoutes : Sets the auth routes
func SetAuthRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/auth/getToken", controllers.GetToken).Methods("POST")
	router.HandleFunc("/api/auth/profile", middleware.AuthMiddleware(controllers.Profile)).Methods("POST")

	return router
}
