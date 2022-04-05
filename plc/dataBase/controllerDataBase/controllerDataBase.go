package controllerDatabase

import (
	"003-API/plc/dataBase/modelDataBase"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var schema modelDataBase.SensorAnalogico

func connectDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=Delfina.0203 dbname=dbPlc port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&schema)
	return db
}

// func ReadDatoS() []model.Dato {
// 	datos := model.Datos
// 	db := connectDb()
// 	db.Find(&datos, "true")
// 	return datos
// }

// func ReadDato(id int) model.Dato {
// 	var dato model.Dato
// 	db := connectDb()
// 	db.First(&dato, id)
// 	return dato
// }

func CreateDato(dato modelDataBase.SensorAnalogico) modelDataBase.SensorAnalogico {
	db := connectDb()
	db.Create(&dato)
	return dato
}

// func UpdateteDato(dato *model.Dato, id int) model.Dato {
// 	var datoAnt model.Dato
// 	datoAct := dato
// 	db := connectDb()
// 	db.First(&datoAnt, id)
// 	datoAnt = *datoAct
// 	db.Save(&datoAnt)
// 	return datoAnt
// }

// func DeleteteDato(id int) model.Dato {
// 	var dato model.Dato
// 	db := connectDb()
// 	db.First(&dato, id)
// 	db.Delete(&model.Datos, id)
// 	return dato
// }
