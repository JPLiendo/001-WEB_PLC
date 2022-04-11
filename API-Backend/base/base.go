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

func ReadDatoS() []model.MachineStates {
	datos := model.Registers
	ConnectDB.Find(&datos, "true")
	return datos
}

func ReadDato(id int) model.MachineStates {
	var dato model.MachineStates

	ConnectDB.First(&dato, id)
	return dato
}

func CreateDato(dato model.MachineStates) model.MachineStates {

	ConnectDB.Create(&dato)
	return dato
}

func UpdateteDato(dato *model.MachineStates) *model.MachineStates {
	ConnectDB.Save(&dato)
	return dato
}

func DeleteteDato(id int) *model.MachineStates {

	ConnectDB.Delete(&model.Registers, id)
	return &model.Registers[id]
}
