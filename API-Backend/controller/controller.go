package controller

import (
<<<<<<< HEAD:API-Backend/controller/controller.go
	"PLC-WEB/API-Backend/base"
	"PLC-WEB/API-Backend/model"
=======
	"003-API/Backend/base"
	"003-API/Backend/model"
>>>>>>> 0d26b8cecef256f1b0ea110e1d851ae7fd3cfe30:Backend/controller/controller.go
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

//Creo la secion global para abrirla una sola vez y poder ejecutar multiples consultas.
func dbConnection() gorm.DB {
	db := base.ConnectDb()
	return *db
}

//GetDatos get a list of Datos
func GetDatos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	db := dbConnection()
	datos := base.ReadDatoS(&db)
	fmt.Fprintln(w, datos)

}

//GetDato get a single of Datos
func GetDato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id_int, _ := strconv.Atoi(vars["id"])
	db := dbConnection()
	dato := base.ReadDato(&db, id_int)
	fmt.Fprintln(w, dato)
}

// //CreateDato create a single of Dato.
func CreateDato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	dato := model.SensorAnalogico{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &dato)
	db := dbConnection()
	dato = base.CreateDato(&db, dato)
	fmt.Fprintln(w, dato)
}

// //UpdateDato update a sigle of Dato.
func UpdateDato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	vars := mux.Vars(r)
	id_int, _ := strconv.Atoi(vars["id"])

	dato := model.SensorAnalogico{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &dato)
	db := dbConnection()
	dato = base.UpdateteDato(&db, &dato, id_int)
	fmt.Fprintln(w, &dato)
}

// //DeleteUsers delete a single of users
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	vars := mux.Vars(r)
	id_int, _ := strconv.Atoi(vars["id"])
	db := dbConnection()
	dato := base.DeleteteDato(&db, id_int)
	fmt.Fprintln(w, dato)
}
