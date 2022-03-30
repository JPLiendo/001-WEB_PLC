package base

import (
	"003-API/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var schema model.Dato

func connectDb() *gorm.DB {
	dsn := "host=localhost user=jpliendo password=Delfina.0203 dbname=db1 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
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

func UpdateteDato(dato *model.Dato) model.Dato {
	datoact := *dato
	db := connectDb()
	db.Save(&datoact)
	return datoact
}
