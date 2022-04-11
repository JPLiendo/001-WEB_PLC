package response

import (
	"PLC-WEB/API-Backend/base"
	"PLC-WEB/API-Backend/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ResponseWr(w http.ResponseWriter, dato *model.MachineStates) {
	err := json.NewEncoder(w).Encode(dato)
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Fprintln(w, "El registro ID:", dato.Id, " no existe, su valor es: ", base.ConnectDB.Error)
	}
}
