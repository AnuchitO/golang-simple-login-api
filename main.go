package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/anuchitprasertsang/golang-login-jwt/token"
)

func main() {
	api := NewAPI(NewRoute())

	log.Fatal(http.ListenAndServe(":8081", api.MakeHandler()))
}

func NewRoute() rest.App {
	router, err := rest.MakeRouter(
		rest.Post("/login", Login),
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

	loginMiddle := &LoginMiddleware{}
	api.Use(loginMiddle)

	api.SetApp(router)
	return
}

func Login(w rest.ResponseWriter, r *rest.Request) {
	body := map[string]string{}

	err := r.DecodeJsonPayload(&body)
	if err != nil {
		w.WriteHeader(400)
		w.WriteJson(err)
	}

	response, status := token.CheckAuthorize(body["user"], body["password"])

	w.WriteHeader(status)
	w.WriteJson(response)
}

type LoginMiddleware struct {
}

func (login *LoginMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {
	return func(w rest.ResponseWriter, r *rest.Request) {
		if r.URL.Path != "/login" {
			err := token.TokenValidator(strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1))
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.WriteJson(map[string]string{"error": err.Error()})
				return
			}
		}
		handler(w, r)
	}
}

func GetUser(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"user": "kob@gmail.com"})
}
