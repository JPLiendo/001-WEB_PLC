package base

import (
	"fmt"

	"github.com/JPLiendo/001-WEB_PLC/API-Backend/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//DataBase stablishes the current database's connection.
var dsn = "host=localhost user=postgres password=12345 dbname=dbPlc port=5432 sslmode=disable TimeZone=Asia/Shanghai"
var ConnectDB = func() (db *gorm.DB) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(`Fail connection database "dbPlc"`)
	} else {
		fmt.Println(`"Succesfull conenction database "dbplc"`)
	}
	return db
}()

//Migrate model to DB.
func MigrateModel() {

	ConnectDB.AutoMigrate(model.MachineStates{})
}

func ReadRegisters() ([]model.MachineStates, *gorm.DB) {
	register := model.Registers
	result := ConnectDB.Find(&register)
	return register, result
}

func ReadRegister(id int) (*model.MachineStates, *gorm.DB) {
	var register model.MachineStates

	result := ConnectDB.First(&register, id)
	return &register, result
}

func CreateRegister(register *model.MachineStates) (*model.MachineStates, *gorm.DB) {

	result := ConnectDB.Create(register)
	return register, result
}

func UpdateRegister(register *model.MachineStates, id int) (*model.MachineStates, *gorm.DB) {
	var oldRegister model.MachineStates
	result := ConnectDB.First(&oldRegister, id)
	oldRegister = *register
	result = ConnectDB.Save(&oldRegister)
	return register, result
}

func DeleteteRegister(id int) (*model.MachineStates, *gorm.DB) {
	var register model.MachineStates
	result := ConnectDB.Delete(&register, id)
	return &register, result
}
