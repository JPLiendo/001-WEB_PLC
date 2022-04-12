package response

import (
	"PLC-WEB/API-Backend/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Response(w http.ResponseWriter, register *model.MachineStates, id int, statusOK, statusNotOk int, message string) {
	if register.Id == id {
		w.WriteHeader(statusOK)
		err := json.NewEncoder(w).Encode(&register)
		if err != nil {
			log.Panic(err)
		}
	} else if register.Id == 0 {
		w.WriteHeader(statusNotOk)
		fmt.Fprintln(w, "Register ID:", id, message)
	}
}
