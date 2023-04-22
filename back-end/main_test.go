package main

import (
	"bytes"
	"encoding/json"
	"github.com/ponsonio/scoir-jcabrera-challenge/back-end/auth"
	"github.com/ponsonio/scoir-jcabrera-challenge/back-end/server"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

var s server.Server

func TestMain(m *testing.M) {
	prov := auth.NewAuthenticatorProvider()
	service := auth.NewAuthenticationService(&prov)
	s = server.NewServer(&service)
	code := m.Run()
	os.Exit(code)
}

func TestCredentialsLogin(t *testing.T) {
	var jsonStr = []byte(`{"user":"scoir", "password": "scoir"}`)
	req, _ := http.NewRequest("POST", "/login/", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["Token"] != token {
		t.Errorf("unexpectect token for user")
	}

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router().ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
