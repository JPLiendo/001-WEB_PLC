package controller

import (
	"003-API/base"
	"003-API/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetDatos get a list of Datos
func GetDatos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	datos := base.ReadDatoS()
	fmt.Fprintln(w, &datos)

}

//GetDato get a single of Datos
func GetDato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id_int, _ := strconv.Atoi(vars["id"])
	dato := base.ReadDato(id_int)
	fmt.Fprintln(w, &dato)
}

// //CreateDato create a single of Dato.
func CreateDato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	dato := model.Dato{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &dato)

	dato = base.CreateDato(dato)
	fmt.Fprintln(w, &dato)
}

// //UpdateDato update a sigle of Dato.
func UpdateDato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id_int, _ := strconv.Atoi(vars["id"])
	dato := base.ReadDato(id_int)
	dato = base.UpdateteDato(&dato)
	fmt.Fprintln(w, &dato)
}

// //DeleteUsers delete a single of users
// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Elimina un usuario existente")
// }
