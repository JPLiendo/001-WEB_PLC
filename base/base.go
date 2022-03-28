package base

import (
	"003-API/model"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var schema model.Dato

func ConnectDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&schema)
	return db
}

func ReadDato(db *gorm.DB) {

	var dato model.Dato
	db.First(&dato)
	fmt.Println(dato)
}
