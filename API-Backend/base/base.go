package base

import (
<<<<<<< HEAD:API-Backend/base/base.go
	"PLC-WEB/API-Backend/model"
=======
	"003-API/Backend/model"
>>>>>>> 0d26b8cecef256f1b0ea110e1d851ae7fd3cfe30:Backend/base/base.go

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var schema model.SensorAnalogico

func ConnectDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=Delfina.0203 dbname=dbPlc port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&schema)
	return db
}

func ReadDatoS(db *gorm.DB) []model.SensorAnalogico {
	datos := model.DatosSensor
	db.Find(&datos, "true")
	return datos
}

func ReadDato(db *gorm.DB, id int) model.SensorAnalogico {
	var dato model.SensorAnalogico

	db.First(&dato, id)
	return dato
}

func CreateDato(db *gorm.DB, dato model.SensorAnalogico) model.SensorAnalogico {

	db.Create(&dato)
	return dato
}

func UpdateteDato(db *gorm.DB, dato *model.SensorAnalogico, id int) model.SensorAnalogico {
	var datoAnt model.SensorAnalogico
	datoAct := dato

	db.First(&datoAnt, id)
	datoAnt = *datoAct
	db.Save(&datoAnt)
	return datoAnt
}

func DeleteteDato(db *gorm.DB, id int) model.SensorAnalogico {
	var dato model.SensorAnalogico

	db.First(&dato, id)
	db.Delete(&model.DatosSensor, id)
	return dato
}
