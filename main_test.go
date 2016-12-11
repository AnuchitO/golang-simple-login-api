package main

import (
	"net/http/httptest"
	"testing"

	"github.com/anuchitprasertsang/golang-simple-login-api/app"
	"github.com/anuchitprasertsang/golang-simple-login-api/routes"
	"github.com/matryer/silk/runner"
)

func TestAPIenpointLogin(t *testing.T) {
	api := app.NewAPI(routes.New())
	s := httptest.NewServer(api.MakeHandler())
	defer s.Close()

	runner.New(t, s.URL).RunFile("./login.silk.md")
}
