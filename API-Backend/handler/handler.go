package handler

import (
	controller "PLC-WEB/API-Backend/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	//mux
	var mux = mux.NewRouter()
	//Handler GetRegisters
	mux.HandleFunc("/register", controller.GetRegisters).Methods(http.MethodGet)
	// //Handler GetRegister
	mux.HandleFunc("/register/{id:[0-9]+}", controller.GetRegister).Methods(http.MethodGet)
	// //Handler CreateRegister
	mux.HandleFunc("/register", controller.CreateRegister).Methods(http.MethodPost)
	// //Handler UpdateRegister
	mux.HandleFunc("/register/{id:[0-9]+}", controller.UpdateRegister).Methods(http.MethodPut)
	//Handler DeleteteRegister
	mux.HandleFunc("/register/{id:[0-9]+}", controller.DeleteteRegister).Methods(http.MethodDelete)
	return mux
}
