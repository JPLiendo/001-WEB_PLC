package controller

import (
	"PLC-WEB/API-Backend/base"
	"PLC-WEB/API-Backend/model"
	"encoding/json"
	"fmt"
	"log"
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
	err := json.NewEncoder(w).Encode(datos)
	if err != nil {
		log.Panic(err)
	}

}

//GetDato get a single of Datos
func GetDato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id_int, _ := strconv.Atoi(vars["id"])
	db := dbConnection()
	dato := base.ReadDato(&db, id_int)
	err := json.NewEncoder(w).Encode(dato)
	if err != nil {
		log.Panic(err)
	}
}

// //CreateDato create a single of Dato.
func CreateDato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	dato := model.SensoresAnalogicos{}
	err := json.NewDecoder(r.Body).Decode(&dato)
	if err != nil {
		log.Panic(err)
	}
	db := dbConnection()

	dato = base.CreateDato(&db, dato)
	err = json.NewEncoder(w).Encode(dato)
	if err != nil {
		log.Panic(err)
	}
}

// //UpdateDato update a sigle of Dato.
func UpdateDato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	vars := mux.Vars(r)
	id_int, _ := strconv.Atoi(vars["id"])
	dato := model.SensoresAnalogicos{}
	db := dbConnection()
	db.First(dato, id_int)

	if dato.Id == id_int {
		dato = base.UpdateteDato(&db, &dato, id_int)
		err := json.NewEncoder(w).Encode(dato)
		if err != nil {
			log.Panic(err)
		}
	} else {
		fmt.Fprintln(w, "El registro ID:", id_int, " no existe, su valor es: ", db.Error)
	}
}

// //DeleteUsers delete a single of users
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	vars := mux.Vars(r)
	id_int, _ := strconv.Atoi(vars["id"])

	dato := model.SensoresAnalogicos{}
	db := dbConnection()
	db.First(dato, id_int)

	if dato.Id == id_int {
		dato = base.DeleteteDato(&db, id_int)
		err := json.NewEncoder(w).Encode(dato)
		if err != nil {
			log.Panic(err)
		}
	} else {
		fmt.Fprintln(w, "El registro ID:", id_int, " no existe, su valor es: ", db.Error)
	}

}
