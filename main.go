package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/anuchitprasertsang/golang-login-jwt/middleware"
	"github.com/anuchitprasertsang/golang-login-jwt/routes"
)

func main() {
	api := NewAPI(routes.New())

	log.Fatal(http.ListenAndServe(":8081", api.MakeHandler()))
}

func NewAPI(router rest.App) (api *rest.Api) {
	api = rest.NewApi()
	api.Use(rest.DefaultProdStack...)

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

	api.Use(middleware.NewLogin())

	api.SetApp(router)
	return
}
