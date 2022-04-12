package handler

import (
	controller "PLC-WEB/API-Backend/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	//mux
	var mux = mux.NewRouter()
	//Handler GetDatos
	mux.HandleFunc("/register", controller.GetRegisters).Methods(http.MethodGet)
	// //Handler GetDato
	mux.HandleFunc("/register/{id:[0-9]+}", controller.GetRegister).Methods(http.MethodGet)
	// //Handler CreateDato
	mux.HandleFunc("/register", controller.CreateRegister).Methods(http.MethodPost)
	// //Handler UpdateDato
	// mux.HandleFunc("/dato/{id:[0-9]+}", controller.UpdateDato).Methods(http.MethodPut)
	//Handler DeleteDato
	//mux.HandleFunc("/dato/{id:[0-9]+}", controller.DeleteUser).Methods(http.MethodDelete)
	return mux
}
