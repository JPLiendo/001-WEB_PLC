package main

import (
	"github.com/JPLiendo/001-WEB_PLC/API-Backend/base"
	"github.com/JPLiendo/001-WEB_PLC/API-Backend/handler"
	"github.com/JPLiendo/001-WEB_PLC/API-Backend/server"
)

func main() {
	//Start connection with DB and migrate model.
	base.MigrateModel()
	//Router
	mux := handler.Router()

	//server
	server.ServerInit(":8080", mux)

}
