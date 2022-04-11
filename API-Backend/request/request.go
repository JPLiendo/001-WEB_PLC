package request

import (
	"PLC-WEB/API-Backend/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetId(r *http.Request) int {
	vars := mux.Vars(r)
	id_int, _ := strconv.Atoi(vars["id"])
	return id_int
}

func GetBody(r *http.Request) *model.MachineStates {
	var dato model.MachineStates
	json.NewDecoder(r.Body).Decode(&dato)
	return &dato
}
