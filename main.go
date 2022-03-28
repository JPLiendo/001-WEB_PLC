package main

import (
	"003-API/base"
	"003-API/handler"
	"003-API/server"
)

func main() {

	//Router
	mux := handler.Router()

	//server
	server.ServerInit(":8080", mux)

	//Base de datos
	db := base.ConnectDb()
	base.ReadDato(db)
}
