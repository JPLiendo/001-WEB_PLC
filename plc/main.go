package main

import (
	plc "003-API/plc/controllerPlc"
	controllerDatabase "003-API/plc/dataBase/controllerDataBase"
	"fmt"
	"time"

	"github.com/robfig/cron"
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
	timeSample(&plc)
}

func createSample(plc *plc.S7) {
	dato := plc.ClientReadDb()
	datoCreado := controllerDatabase.CreateDato(dato)
	fmt.Println(datoCreado)

}
func timeSample(plc *plc.S7) {
	c := cron.New()
	defer c.Stop()
	c.AddFunc("@every 10s", func() { createSample(plc) })
	c.Start()
	select {}
}
