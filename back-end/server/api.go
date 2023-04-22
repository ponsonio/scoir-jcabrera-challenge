package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ponsonio/scoir-jcabrera-challenge/back-end/auth"
	"io"
	"net/http"
)

type api struct {
	router                http.Handler
	authenticationService *auth.AuthenticationService
}

type Server interface {
	Router() http.Handler
}

func NewServer(authenticationService *auth.AuthenticationService) Server {
	a := &api{
		authenticationService: authenticationService,
	}

	r := mux.NewRouter()

	r.HandleFunc("/login/", a.loginWithCredentials).Methods(http.MethodPost)

	a.router = r
	return a
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func (a *api) loginWithCredentials(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	var userLogin auth.UserLogin
	reqBody, _ := io.ReadAll(r.Body)

	marErr := json.Unmarshal(reqBody, &userLogin)

	if marErr != nil {
		http.Error(w, marErr.Error(), http.StatusInternalServerError)
		return
	}
	service := *a.authenticationService
	token, authErr := service.AuthenticateWithLogin(userLogin)

	if authErr != nil {
		http.Error(w, authErr.Error(), http.StatusUnauthorized)
		return
	}

	resErr := json.NewEncoder(w).Encode(token)
	if resErr != nil {
		http.Error(w, resErr.Error(), http.StatusInternalServerError)
		return
	}

}

func (a *api) Router() http.Handler {
	return a.router
}
