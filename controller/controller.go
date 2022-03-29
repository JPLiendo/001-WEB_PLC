package controller

import (
	"003-API/base"
	"003-API/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// //GetUsers get a list of users
// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "GetUser devuelve una lista de usuarios")
// }

//GetUser get a single of users
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	id_str := r.URL.Query().Get("id")
	id_int, _ := strconv.Atoi(id_str)
	base.ReadDato(w, id_int)
}

// //CreateUsers create a single of users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	dato := model.Dato{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &dato)

	base.CreateDato(w, dato)
}

// //UpdateUsers update a sigle of users
// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Edita un usuario existente")
// }

// //DeleteUsers delete a single of users
// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Elimina un usuario existente")
// }
