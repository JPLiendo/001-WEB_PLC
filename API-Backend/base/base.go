package base

import (
	"PLC-WEB/API-Backend/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//DataBase stablishes the current database's connection.
var dsn = "host=localhost user=postgres password=Delfina.0203 dbname=dbPlc port=5432 sslmode=disable TimeZone=Asia/Shanghai"
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
	registers := model.Registers
	result := ConnectDB.Find(&registers, "true")
	return registers, result
}

func ReadRegister(id int) (model.MachineStates, *gorm.DB) {
	var register model.MachineStates
	result := ConnectDB.First(&register, id)
	return register, result
}

func CreateRegister(register model.MachineStates) (model.MachineStates, *gorm.DB) {
	result := ConnectDB.First(&register, register.Id)
	if result.RowsAffected != 0 {
		result = ConnectDB.Create(&register)

	}
	return register, result
}

// func UpdateteDato(dato *model.MachineStates) *model.MachineStates {
// 	ConnectDB.Save(&dato)
// 	return dato
// }

// func DeleteteDato(id int) *model.MachineStates {

// 	ConnectDB.Delete(&model.Registers, id)
// 	return &model.Registers[id]
// }
