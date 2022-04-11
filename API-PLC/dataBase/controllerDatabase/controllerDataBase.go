package controllerDataBase

import (
	connectionDatabase "PLC-WEB/API-PLC/dataBase/connectionDataBase"
	"log"

	"gorm.io/gorm"
)

//Type MachineState.
type MachineStates struct {
	//Id             uint `gorm:"primaryKey"`
	Running        bool
	AlertStop      bool
	FuncionalStop  bool
	StopByFoward   bool
	StopByBackward bool
	UnderSpeed     bool
}

func MigrateModel() {
	if err := connectionDatabase.ConnectDB.AutoMigrate(MachineStates{}); err != nil {
		log.Panic("Model migration fail ", err)
	}
}

//Interface.
type Machine interface {
	createRegister(*gorm.DB)
}

//CreateRegister method create a new register.
func (f *MachineStates) CreateRegister(db *gorm.DB) {
	db.Create(f)
}
