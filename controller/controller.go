package controller

import (
	"fmt"
	"net/http"
)

//GetUsers get a list of users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GetUser devuelve una lista de usuarios")
}

//GetUser get a single of users
func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Devuelve un usuario")
}

//CreateUsers create a single of users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, `{"usuario":"Crea un nuevo usuario"}`)
}

//UpdateUsers update a sigle of users
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Edita un usuario existente")
}

//DeleteUsers delete a single of users
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Elimina un usuario existente")
}
