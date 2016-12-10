package main

import (
	"net/http/httptest"
	"testing"

	"github.com/matryer/silk/runner"
)

func TestAPIenpointLogin(t *testing.T) {
	api := NewAPI(NewRoute())
	s := httptest.NewServer(api.MakeHandler())
	defer s.Close()

	runner.New(t, s.URL).RunFile("./login.silk.md")
}
