package base

import (
	"003-API/Backend/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var schema model.SensorAnalogico

func connectDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=Delfina.0203 dbname=dbPlc port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&schema)
	return db
}

func ReadDatoS() []model.SensorAnalogico {
	datos := model.DatosSensor
	db := connectDb()
	db.Find(&datos, "true")
	return datos
}

func ReadDato(id int) model.SensorAnalogico {
	var dato model.SensorAnalogico
	db := connectDb()
	db.First(&dato, id)
	return dato
}

func CreateDato(dato model.SensorAnalogico) model.SensorAnalogico {
	db := connectDb()
	db.Create(&dato)
	return dato
}

func UpdateteDato(dato *model.SensorAnalogico, id int) model.SensorAnalogico {
	var datoAnt model.SensorAnalogico
	datoAct := dato
	db := connectDb()
	db.First(&datoAnt, id)
	datoAnt = *datoAct
	db.Save(&datoAnt)
	return datoAnt
}

func DeleteteDato(id int) model.SensorAnalogico {
	var dato model.SensorAnalogico
	db := connectDb()
	db.First(&dato, id)
	db.Delete(&model.DatosSensor, id)
	return dato
}
