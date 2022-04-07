package main

import (
	controllerDatabase "PLC-WEB/API-PLC/database/controllerDataBase"
	controllerPlc "PLC-WEB/API-PLC/plc"
	"fmt"
	"time"

	"github.com/robfig/cron"
	"gorm.io/gorm"
)

func main() {
	plc := controllerPlc.S7{
		S7Connection: controllerPlc.S7Connection{
			Addr:        "192.168.0.10",
			Rack:        0,
			Slot:        2,
			TimeOut:     200 * time.Second,
			IdleTimeout: 200 * time.Second,
		},
		S7Data: controllerPlc.S7Data{
			S7DataRead: controllerPlc.S7DataRead{
				DbNumber: 400,
				Start:    0,
				Size:     32,
				Buffer:   make([]byte, 32),
			},
		},
	}
	connectDb := connectDb()
	timeSample(&connectDb, &plc)
}

//connectDb devuelve el puntero a la conexión a la DB.
func connectDb() gorm.DB {
	db := controllerDatabase.ConnectDb()
	return *db
}

//createRegister crea un nuevo registro en el DB.
func createRegister(connectDb *gorm.DB, plc *controllerPlc.S7) {
	registro := plc.ClientReadDb()
	newRegister := controllerDatabase.CreateRegister(connectDb, registro)
	fmt.Println(newRegister)

}

//timeSample ejecuta la consulta a la DB en una frecuencia definida.
func timeSample(connectDb *gorm.DB, plc *controllerPlc.S7) {
	c := cron.New()
	defer c.Stop()
	c.AddFunc("@every 10s", func() { createRegister(connectDb, plc) })
	c.Start()
	select {}
}
