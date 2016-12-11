package middleware

import (
	"net/http"
	"strings"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/anuchitprasertsang/golang-simple-login-api/token"
)

var TokenStack = []rest.Middleware{
	&Token{},
}

type Token struct {
}

func (t *Token) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {
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
