package controller

import (
	"PLC-WEB/API-Backend/base"
	"PLC-WEB/API-Backend/model"
	"PLC-WEB/API-Backend/request"
	"PLC-WEB/API-Backend/response"
	"encoding/json"
	"log"
	"net/http"
)

//GetDatos get a list of Datos
func GetDatos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	datos := base.ReadDatoS()
	err := json.NewEncoder(w).Encode(datos)
	if err != nil {
		log.Panic(err)
	}

}

//GetDato get a single of Datos
func GetDato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	id := request.GetId(r)
	dato := base.ReadDato(id)
	err := json.NewEncoder(w).Encode(dato)
	if err != nil {
		log.Panic(err)
	}
}

// //CreateDato create a single of Dato.
func CreateDato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	dato := model.MachineStates{}
	err := json.NewDecoder(r.Body).Decode(&dato)
	if err != nil {
		log.Panic(err)
	}

	dato = base.CreateDato(dato)
	err = json.NewEncoder(w).Encode(dato)
	if err != nil {
		log.Panic(err)
	}
}

// //UpdateDato update a sigle of Dato.
func UpdateDato(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusAccepted)

	Id := request.GetId(r)
	dato := base.ReadDato(Id)

	if dato.Id == Id {
		dato = *request.GetBody(r)
		base.UpdateteDato(&dato)

	}
	response.ResponseWr(w, &dato)
}

// //DeleteUsers delete a single of users
// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusAccepted)

// 	vars := mux.Vars(r)
// 	id_int, _ := strconv.Atoi(vars["id"])

// 	dato := model.MachineStates{}

// 	base.ConnectDB.First(dato, id_int)

// 	if int(dato.Id) == id_int {
// 		dato = base.DeleteteDato(id_int)
// 		err := json.NewEncoder(w).Encode(dato)
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 	} else {
// 		fmt.Fprintln(w, "El registro ID:", id_int, " no existe, su valor es: ", base.ConnectDB.Error)
// 	}

// }
