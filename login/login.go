package login

import "github.com/ant0ine/go-json-rest/rest"

func Login(w rest.ResponseWriter, r *rest.Request) {
	body := map[string]string{}

	err := r.DecodeJsonPayload(&body)
	if err != nil {
		w.WriteHeader(400)
		w.WriteJson(err)
	}

	response, status := CheckAuthenticate(body["user"], body["password"])

	w.WriteHeader(status)
	w.WriteJson(response)
}
