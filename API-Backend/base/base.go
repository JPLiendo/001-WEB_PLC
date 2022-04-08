package base

import (
	"PLC-WEB/API-Backend/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var schema model.SensoresAnalogicos

func ConnectDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=Delfina.0203 dbname=dbPlc port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&schema)
	return db
}

func ReadDatoS(db *gorm.DB) []model.SensoresAnalogicos {
	datos := model.DatosSensor
	db.Find(&datos, "true")
	return datos
}

func ReadDato(db *gorm.DB, id int) model.SensoresAnalogicos {
	var dato model.SensoresAnalogicos

	db.First(&dato, id)
	return dato
}

func CreateDato(db *gorm.DB, dato model.SensoresAnalogicos) model.SensoresAnalogicos {

	db.Create(&dato)
	return dato
}

func UpdateteDato(db *gorm.DB, dato *model.SensoresAnalogicos, id int) model.SensoresAnalogicos {
	var datoAnt model.SensoresAnalogicos
	datoAct := dato

	db.First(&datoAnt, id)
	datoAnt = *datoAct
	db.Save(&datoAnt)
	return datoAnt
}

func DeleteteDato(db *gorm.DB, id int) model.SensoresAnalogicos {
	var dato model.SensoresAnalogicos

	db.First(&dato, id)
	db.Delete(&model.DatosSensor, id)
	return dato
}
