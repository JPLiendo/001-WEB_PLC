package main

import (
	"PLC-WEB/Backend/handler"
	"PLC-WEB/Backend/server"
)

func main() {

	//Router
	mux := handler.Router()

	//server
	server.ServerInit(":8080", mux)

}
