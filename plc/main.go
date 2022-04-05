package main

import (
	plc "PLC-WEB/plc/controllerPlc"
	controllerDatabase "PLC-WEB/plc/dataBase/controllerDataBase"
	"fmt"
	"time"
)

func main() {
	plc := plc.S7{
		S7Connection: plc.S7Connection{
			Addr:        "192.168.0.1",
			Rack:        0,
			Slot:        2,
			TimeOut:     200 * time.Second,
			IdleTimeout: 200 * time.Second,
		},
		S7Data: plc.S7Data{
			S7DataRead: plc.S7DataRead{
				DbNumber: 1,
				Start:    0,
				Size:     32,
				Buffer:   make([]byte, 32),
			},
		},
	}

	dato := plc.ClientReadDb()
	datoCreado := controllerDatabase.CreateDato(dato)
	fmt.Println(datoCreado)

}
