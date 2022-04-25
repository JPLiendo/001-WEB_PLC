package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/JPLiendo/001-WEB_PLC/API-Backend/base"
	"github.com/JPLiendo/001-WEB_PLC/API-Backend/model"
	"github.com/JPLiendo/001-WEB_PLC/API-Backend/request"
)

//GetRegisters get a list of registers.
func GetRegisters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	registers, _ := base.ReadRegisters()
	err := json.NewEncoder(w).Encode(&registers)
	if err != nil {
		log.Panic(err)
	}

}

//GetRegister get a single of register.
func GetRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := request.GetId(r)
	register, result := base.ReadRegister(id)

	if register.Id == id {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(&register)
		if err != nil {
			log.Panic(err)
		}
	} else if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Register ", id, "not Found")
	}

}

//CreateRegister create a single of register.
func CreateRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	register := model.MachineStates{}
	err := json.NewDecoder(r.Body).Decode(&register)
	if err != nil {
		log.Panic(err)
	}
	_, result := base.CreateRegister(&register)

	if result.RowsAffected >= 1 {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Register ", register.Id, "created")
	} else if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Register ", register.Id, "not crated")

	}
}

// //UpdateRegister update a sigle of register.
func UpdateRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	register := model.MachineStates{}
	err := json.NewDecoder(r.Body).Decode(&register)
	if err != nil {
		log.Panic(err)
	}
	_, result := base.UpdateRegister(&register, register.Id)

	if result.RowsAffected >= 1 {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintln(w, "Register ", register.Id, "updated")
	} else if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Register ", register.Id, "not found")

	}
}

//DeleteteRegister delete a single of register.
func DeleteteRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := request.GetId(r)
	_, result := base.DeleteteRegister(id)

	if result.RowsAffected >= 1 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Register ", id, "deleted.")
	} else if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Register ", id, "not Found")
	}
}
