package token

func CheckAuthorize(user, passwd string) (map[string]string, int) {
	resp := map[string]string{}

	if user == "kob@gmail.com" && passwd == "aobaob" {
		resp["token"] = CreateToken(user)
		return resp, 200
	}

	resp["error"] = "user or password wrong"
	return resp, 401
}
