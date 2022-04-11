package main

import (
	connectionDatabase "PLC-WEB/API-PLC/dataBase/connectionDataBase"
	controllerDataBase "PLC-WEB/API-PLC/dataBase/controllerDatabase"

	"PLC-WEB/API-PLC/plc/connectionPlc"
	"PLC-WEB/API-PLC/plc/controllerPlc"

	"github.com/robfig/cron"
	"gorm.io/gorm"
)

func main() {

	controllerDataBase.MigrateModel()
	var fillerstates controllerDataBase.MachineStates
	timeSample(&fillerstates, connectionDatabase.ConnectDB)
}

func Register(machineStates *controllerDataBase.MachineStates, newRegister *gorm.DB) {
	machineStates = controllerPlc.ReadBoolByte(&connectionPlc.PLC, machineStates)
	machineStates.CreateRegister(newRegister)

}

//timeSample ejecuta la consulta a la DB en una frecuencia definida.
func timeSample(machineStates *controllerDataBase.MachineStates, newRegister *gorm.DB) {
	c := cron.New()
	defer c.Stop()
	c.AddFunc("@every 10s", func() { Register(machineStates, newRegister) })
	c.Start()
	select {}
}
