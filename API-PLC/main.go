package main

import (
	connectionDatabase "PLC-WEB/API-PLC/dataBase/connectionDataBase"
	controllerDatabase "PLC-WEB/API-PLC/dataBase/controllerDataBase"

	"github.com/robfig/cron"
	"gorm.io/gorm"
)

func main() {
	var fillerstates = controllerDatabase.MachineStates{}
	connectionDatabase.DataBase.AutoMigrate(fillerstates)
	timeSample(&fillerstates, connectionDatabase.DataBase)
}

//timeSample ejecuta la consulta a la DB en una frecuencia definida.
func timeSample(machineStates *controllerDatabase.MachineStates, newRegister *gorm.DB) {
	c := cron.New()
	defer c.Stop()
	c.AddFunc("@every 10s", func() { controllerDatabase.CreateMachineRegister(machineStates, newRegister) })
	c.Start()
	select {}
}
