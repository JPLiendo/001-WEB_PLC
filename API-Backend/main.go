package main

import (
	"003-API/Backend/handler"
	"003-API/Backend/server"
)

func main() {

	//Router
	mux := handler.Router()

	//server
	server.ServerInit(":8080", mux)

}
