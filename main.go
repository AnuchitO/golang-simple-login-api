package main

import (
	"log"
	"net/http"

	"github.com/anuchitprasertsang/golang-login-jwt/app"
	"github.com/anuchitprasertsang/golang-login-jwt/routes"
)

func main() {
	api := app.NewAPI(routes.New())

	log.Fatal(http.ListenAndServe(":8081", api.MakeHandler()))
}
