package handler

import (
	"net/http"

	c "003-API/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	//mux
	var mux = mux.NewRouter()
	//Handler GetDatos
	mux.HandleFunc("/dato", c.GetDatos).Methods(http.MethodGet)
	//Handler GetDato
	mux.HandleFunc("/dato/{id:[0-9]+}", c.GetDato).Methods(http.MethodGet)
	//Handler CreateDato
	mux.HandleFunc("/dato", c.CreateDato).Methods(http.MethodPost)
	//Handler UpdateDato
	mux.HandleFunc("/dato/{id:[0-9]+}", c.UpdateDato).Methods(http.MethodPut)
	//Handler DeleteDato
	//mux.HandleFunc("/Dato/{id:[0-9]+}", c.DeleteUser).Methods(http.MethodDelete)
	return mux
}
