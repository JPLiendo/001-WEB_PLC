package controllerDatabase

import (
	"PLC-WEB/API-PLC/dataBase/modelDataBase"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Genero el schema de la BD a partir del modelo.
var schema modelDataBase.SensoresAnologicos

/*ConnetDB devuelve el puntero de la conexión a la DB, es exportada para poder ser llamada desde main
y generar una sola sesion, con esto se pueden realizar multiples consultas desde la misma sesión*/
func ConnectDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=Delfina.0203 dbname=dbPlc port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(`No se ha podido conectarse con la base de datos "dbPlc"`)
	}
	db.AutoMigrate(&schema)
	return db
}

/*CreateRegister crea un nuevo registro en la DB*/
func CreateRegister(db *gorm.DB, dato *modelDataBase.SensoresAnologicos) *modelDataBase.SensoresAnologicos {
	db.Create(dato)
	return dato
}
