package main

import (
	"fmt"
	"kobapp/customer"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

type LoginMiddleware struct {
}

func (login *LoginMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {
	return func(w rest.ResponseWriter, r *rest.Request) {
		fmt.Println("before execute handler")

		if r.URL.Path != "/login" {
			token := r.Header.Get("Authorization")
			if token != "1234567890" {
				w.WriteHeader(401)
				w.WriteJson(map[string]string{"error": "permission denine"})
				return
			}
		}

		handler(w, r)
		fmt.Println("after execute handler")
	}
}

func main() {
	api := rest.NewApi()
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

	router, err := rest.MakeRouter(
		rest.Post("/login", Login),
		rest.Post("/customers", customer.CustomerAPI),
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)

	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func Login(w rest.ResponseWriter, r *rest.Request) {
	body := map[string]string{}

	err := r.DecodeJsonPayload(&body)
	if err != nil {
		w.WriteHeader(400)
		w.WriteJson(err)
	}

	response, status := checkAuthorize(body["user"], body["password"])

	w.WriteHeader(status)
	w.WriteJson(response)
}

func checkAuthorize(user, passwd string) (map[string]string, int) {
	resp := map[string]string{}

	if user == "kob" && passwd == "1234" {
		resp["token"] = "1234567890"
		return resp, 200
	}

	resp["error"] = "user or password wrong"
	return resp, 401
}
