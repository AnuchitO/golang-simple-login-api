package routes

import (
	"log"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/anuchitprasertsang/golang-simple-login-api/customers"
	"github.com/anuchitprasertsang/golang-simple-login-api/login"
)

func New() rest.App {
	router, err := rest.MakeRouter(
		rest.Post("/login", login.Login),
		rest.Get("/customers", customers.CustomerAPI),
	)

	if err != nil {
		log.Fatal(err)
	}

	return router
}
