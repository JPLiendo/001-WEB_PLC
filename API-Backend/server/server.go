package server

import (
	"log"
	"net/http"
)

func ServerInit(Host string, handler http.Handler) {

	error := http.ListenAndServe(Host, handler)
	if error != nil {
		log.Fatal("Fallo de conección al servidor!!!", error)
	} else {
		log.Println("Conección exitosa con el servidor")
	}
}
