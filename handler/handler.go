package handler

import (
	"net/http"

	c "003-API/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	//mux
	var mux = mux.NewRouter()
	//Handler GetUsers
	mux.HandleFunc("/user", c.GetUsers).Methods(http.MethodGet)
	//Handler GetUser
	mux.HandleFunc("/user/{id:[0-9]+}", c.GetUser).Methods(http.MethodGet)
	//Handler CreateUser
	mux.HandleFunc("/user", c.CreateUser).Methods(http.MethodPost)
	//Handler UpdateUser
	mux.HandleFunc("/user/{id:[0-9]+}", c.UpdateUser).Methods(http.MethodPut)
	//Handler DeleteUser
	mux.HandleFunc("/user/{id:[0-9]+}", c.DeleteUser).Methods(http.MethodDelete)
	return mux
}
