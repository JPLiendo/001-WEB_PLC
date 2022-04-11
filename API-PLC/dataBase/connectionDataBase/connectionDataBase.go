package connectionDatabase

import (
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
