package plc

import (
	"fmt"
	"log"
	"time"

	"github.com/robinson/gos7"
)

//Type s7Connction.
type S7Connection struct {
	Addr        string
	Rack        int
	Slot        int
	TimeOut     time.Duration
	IdleTimeout time.Duration
}

type S7DataRead struct {
	DbNumber int
	Start    int
	Size     int
	Buffer   []byte
}

// type S7DataWrite struct {
// 	DbNumber int
// 	Start    int
// 	Size     int
// 	Buffer   []byte
// }

type S7Data struct {
	S7DataRead
	// S7DataWrite
}
type S7 struct {
	S7Connection
	S7Data
}

func handlerConnection(s7 *S7Connection) *gos7.TCPClientHandler {
	handlerConnection := gos7.NewTCPClientHandler(s7.Addr, s7.Rack, s7.Slot)
	handlerConnection.Timeout = s7.TimeOut
	handlerConnection.IdleTimeout = s7.IdleTimeout
	//handlerConnection.Logger = log.New(os.Stdout, "tcp: ", log.LstdFlags)
	return handlerConnection
}

type S7Connction interface {
	ClientReadDb()
}

//Comienzo de las funciones CRUD para el PLC.

//Client reads a complete DB by number.
func (s7 *S7) ClientReadDb() {
	handler := handlerConnection(&s7.S7Connection)
	err := handler.Connect()
	if err != nil {
		log.Panic("No se pudo conectar con el PLC")
	}
	defer handler.Close()
	client := gos7.NewClient(handler)
	err = client.AGReadDB(s7.DbNumber, s7.Start, s7.Size, s7.Buffer)
	if err != nil {
		log.Panic(err)
	}
	var s7Helper gos7.Helper
	var result float32
	for i := 0; i <= len(s7.Buffer)-1; i = i + 4 {
		s7Helper.GetValueAt(s7.Buffer, i, &result)
		fmt.Println(result)

	}

}
