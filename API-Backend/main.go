package main

import (
	"PLC-WEB/API-Backend/handler"
	"PLC-WEB/API-Backend/server"
)

func main() {

	//Router
	mux := handler.Router()

	//server
	server.ServerInit(":8080", mux)

}
