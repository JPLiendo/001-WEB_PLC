package base

import (
	"PLC-WEB/Backend/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var schema model.Dato

func connectDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=Delfina.0203 dbname=db1 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&schema)
	return db
}

func ReadDatoS() []model.Dato {
	datos := model.Datos
	db := connectDb()
	db.Find(&datos, "true")
	return datos
}

func ReadDato(id int) model.Dato {
	var dato model.Dato
	db := connectDb()
	db.First(&dato, id)
	return dato
}

func CreateDato(dato model.Dato) model.Dato {
	db := connectDb()
	db.Create(&dato)
	return dato
}

func UpdateteDato(dato *model.Dato, id int) model.Dato {
	var datoAnt model.Dato
	datoAct := dato
	db := connectDb()
	db.First(&datoAnt, id)
	datoAnt = *datoAct
	db.Save(&datoAnt)
	return datoAnt
}

func DeleteteDato(id int) model.Dato {
	var dato model.Dato
	db := connectDb()
	db.First(&dato, id)
	db.Delete(&model.Datos, id)
	return dato
}
