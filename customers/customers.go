package customers

import "github.com/ant0ine/go-json-rest/rest"

func CustomerAPI(w rest.ResponseWriter, r *rest.Request) {
	customer := []string{
		"apichat",
		"anuchit",
		"wuttinun",
	}

	w.WriteJson(customer)
}
