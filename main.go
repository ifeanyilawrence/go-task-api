package main

import (
	"log"
	"net/http"

	"github.com/ifeanyilawrence/go-task-api/routers"
)

func main() {
	hndlr := routers.InitRoutes()

	log.Fatal(http.ListenAndServe(":8080", hndlr))
}
