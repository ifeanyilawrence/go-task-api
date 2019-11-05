package routers

import "github.com/gorilla/mux"

//InitRoutes : initializes all routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetTaskRoutes(router)
	router = SetUserRoutes(router)
	return router
}
