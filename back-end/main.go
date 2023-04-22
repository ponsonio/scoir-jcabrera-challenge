package main

import (
	"github.com/ponsonio/scoir-jcabrera-challenge/back-end/auth"
	"github.com/ponsonio/scoir-jcabrera-challenge/back-end/server"
	"log"
	"net/http"
)

func main() {
	//this allows to inject a mock or replace the implementation
	prov := auth.NewAuthenticatorProvider()
	service := auth.NewAuthenticationService(&prov)
	server := server.NewServer(&service)
	log.Fatal(http.ListenAndServe(":8080", server.Router()))
}
