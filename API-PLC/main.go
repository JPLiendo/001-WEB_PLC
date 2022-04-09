package main

import (
	controllerDatabase "PLC-WEB/API-PLC/database/controllerDataBase"
	"PLC-WEB/API-PLC/plc/connectionPlc"
	controllerPlc "PLC-WEB/API-PLC/plc/controller.plc"
	"fmt"
)

func main() {
	fillerData := controllerDatabase.FillerStates{}
	newRegister := controllerPlc.SaveReadDb(&connectionPlc.PLC, fillerData)
	fmt.Println(newRegister)

}

// //timeSample ejecuta la consulta a la DB en una frecuencia definida.
// func timeSample(connectDb *gorm.DB, plc *controllerPlc.S7) {
// 	newRegister:=
// 	c := cron.New()
// 	defer c.Stop()
// 	c.AddFunc("@every 10s", func() { controllerDatabase.CreateMachineRegister() })
// 	c.Start()
// 	select {}
// }
