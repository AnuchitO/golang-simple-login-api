package token

import "net/http"

func CheckAuthorize(user, passwd string) (map[string]string, int) {
	resp := map[string]string{}

	if user == "kob@gmail.com" && passwd == "aobaob" {
		resp["token"] = CreateToken(user)
		return resp, http.StatusOK
	}

	resp["error"] = "user or password wrong"
	return resp, http.StatusUnauthorized
}
