package main

import (
	"log"
	"net/http"

	"github.com/anuchitprasertsang/golang-simple-login-api/app"
	"github.com/anuchitprasertsang/golang-simple-login-api/routes"
)

func main() {
	api := app.NewAPI(routes.New())

	log.Fatal(http.ListenAndServe(":8081", api.MakeHandler()))
}
