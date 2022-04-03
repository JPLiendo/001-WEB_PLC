package plc

import "github.com/robinson/gos7"

type s7 struct {
	addr string
	rack int
	slot int
}


func (s7 *s7) clientPLC() gos7.Client {
	handlerClient := gos7.NewTCPClientHandler(s7.addr, s7.rack, s7.slot)
	client := gos7.NewClient(handlerClient)
	return client
}

func (s7 *s7) connectPlc(){
	
}


//Comienzo de las funciones CRUD para el PLC.

//ReadPlcDb reads a complete DB by number.
func ClientReadDb(client gos7.Client) {
	client.


}

//WritePlcDb write a complete DB by number.

func writePlcDb() {}
