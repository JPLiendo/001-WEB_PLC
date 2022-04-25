package request

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JPLiendo/001-WEB_PLC/API-Backend/model"

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
