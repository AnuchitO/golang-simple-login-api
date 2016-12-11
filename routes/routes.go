package routes

import (
	"log"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/anuchitprasertsang/golang-login-jwt/customer"
	"github.com/anuchitprasertsang/golang-login-jwt/login"
)

func New() rest.App {
	router, err := rest.MakeRouter(
		rest.Post("/login", login.Login),
		rest.Get("/customers", customer.CustomerAPI),
	)

	if err != nil {
		log.Fatal(err)
	}

	return router
}
