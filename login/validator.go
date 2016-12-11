package login

import (
	"net/http"

	"github.com/anuchitprasertsang/golang-login-jwt/token"
)

func CheckAuthenticate(user, passwd string) (map[string]string, int) {
	resp := map[string]string{}

	if has(user, passwd) {
		resp["token"] = token.CreateToken(user)
		return resp, http.StatusOK
	}

	resp["error"] = "user or password wrong"
	return resp, http.StatusUnauthorized
}

func has(user, passwd string) bool {
	return user == "kob@gmail.com" && passwd == "aobaob"
}
