package main

import (
	"PLC-WEB/API-Backend/base"
	"PLC-WEB/API-Backend/handler"
	"PLC-WEB/API-Backend/server"
)

func main() {
	//Start connection with DB and migrate model.
	base.MigrateModel()
	//Router
	mux := handler.Router()

	//server
	server.ServerInit(":8080", mux)

}
