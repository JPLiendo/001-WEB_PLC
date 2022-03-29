package base

import (
	"003-API/model"
	"fmt"
	"net/http"

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

func ReadDato(w http.ResponseWriter, id int) {
	var dato model.Dato
	db := connectDb()
	db.First(&dato, id)
	fmt.Fprintln(w, &dato)
}

func CreateDato(w http.ResponseWriter, dato model.Dato) {
	db := connectDb()
	db.Create(&dato)
	fmt.Fprintln(w, &dato)
}
