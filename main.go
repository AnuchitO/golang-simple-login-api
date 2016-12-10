package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/anuchitprasertsang/golang-login-jwt/login"
	"github.com/anuchitprasertsang/golang-login-jwt/middleware"
)

func main() {
	api := NewAPI(NewRoute())

	log.Fatal(http.ListenAndServe(":8081", api.MakeHandler()))
}

func NewRoute() rest.App {
	router, err := rest.MakeRouter(
		rest.Post("/login", login.Login),
		rest.Get("/users", GetUser),
	)

	if err != nil {
		log.Fatal(err)
	}
	return router
}

func NewAPI(router rest.App) (api *rest.Api) {
	api = rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	allowedMethods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	allowedHeaders := []string{
		"Accept",
		"Authorization",
		"X-Real-IP",
		"Content-Type",
		"X-Custom-Header",
		"Language",
		"Origin",
	}
	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return true
		},
		AllowedMethods:                allowedMethods,
		AllowedHeaders:                allowedHeaders,
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})

	loginMiddle := &middleware.LoginMiddleware{}
	api.Use(loginMiddle)

	api.SetApp(router)
	return
}

func GetUser(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"user": "kob@gmail.com"})
}
