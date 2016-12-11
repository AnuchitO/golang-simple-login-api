package routes

import (
	"log"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/anuchitprasertsang/golang-login-jwt/login"
)

func New() rest.App {
	router, err := rest.MakeRouter(
		rest.Post("/login", login.Login),
		rest.Get("/users", GetUser),
	)

	if err != nil {
		log.Fatal(err)
	}

	return router
}

func GetUser(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"user": "kob@gmail.com"})
}
