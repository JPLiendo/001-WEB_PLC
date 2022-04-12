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

//GetRegisters get a list of registers.
func GetRegisters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	registers, _ := base.ReadRegisters()
	err := json.NewEncoder(w).Encode(registers)
	if err != nil {
		log.Panic(err)
	}

}

//GetRegister get a single of registers.
func GetRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := request.GetId(r)
	register, _ := base.ReadRegister(id)
	response.Response(w, &register, id, http.StatusAccepted, http.StatusNotFound, "does not exist.")
}

//CreateRegister create a single of register.
func CreateRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	register := model.MachineStates{}
	err := json.NewDecoder(r.Body).Decode(&register)
	if err != nil {
		log.Panic(err)
	}
	register, _ = base.CreateRegister(register)
	response.Response(w, &register, register.Id, http.StatusCreated, http.StatusBadRequest, "already exist.")
}

// // //UpdateDato update a sigle of Dato.
// func UpdateDato(w http.ResponseWriter, r *http.Request) {
// 	// w.Header().Set("Content-Type", "application/json")
// 	// w.WriteHeader(http.StatusAccepted)

// 	Id := request.GetId(r)
// 	dato := base.ReadDato(Id)

// 	if dato.Id == Id {
// 		dato = *request.GetBody(r)
// 		base.UpdateteDato(&dato)

// 	}
// 	response.ResponseWr(w, &dato)
// }

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
