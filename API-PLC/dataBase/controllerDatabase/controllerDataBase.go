package controllerDataBase

import (
	"log"

	connectionDatabase "github.com/JPLiendo/001-WEB_PLC/API-PLC/dataBase/connectionDataBase"

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
