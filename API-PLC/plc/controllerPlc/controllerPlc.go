package controllerPlc

import (
	"fmt"

	controllerDataBase "github.com/JPLiendo/001-WEB_PLC/API-PLC/dataBase/controllerDatabase"

	"log"
	"time"

	"github.com/robinson/gos7"
)

//Type S7Connection sets the connection's parameters.
type S7Connection struct {
	Addr        string
	Rack        int
	Slot        int
	TimeOut     time.Duration
	IdleTimeout time.Duration
}

//Type S7DataRead sets the data read's parameters.
type S7DataRead struct {
	DbNumber int
	Start    int
	Size     int
	Buffer   []byte
}

//Type S7DataWrite sets the data write's parameters.
// Tipo S7DataWrite struct {
// 	DbNumber int
// 	Start    int
// 	Size     int
// 	Buffer   []byte
// }

//Type S7Data cluster read/wite data.
type S7Data struct {
	S7DataRead
	// S7DataWrite
}

//Type S7 cluster all data from the PLC.
type S7 struct {
	S7Connection
	S7Data
}

//handlerConnection create the handler to set the connection parameters.
func handlerConnection(s7 *S7Connection) *gos7.TCPClientHandler {
	handlerConnection := gos7.NewTCPClientHandler(s7.Addr, s7.Rack, s7.Slot)
	handlerConnection.Timeout = s7.TimeOut
	handlerConnection.IdleTimeout = s7.IdleTimeout
	//handlerConnection.Logger = log.New(os.Stdout, "tcp: ", log.LstdFlags)
	return handlerConnection
}

//clientReadDb reads a single DB according with data type given.
func (s7 *S7) clientReadDb() []byte {
	handler := handlerConnection(&s7.S7Connection)
	if err := handler.Connect(); err != nil {
		log.Panic("PLC connection fail, ip: ", handler.Address, "\n", err)
	}
	defer handler.Close()

	client := gos7.NewClient(handler)
	if err := client.AGReadDB(s7.DbNumber, s7.Start, s7.Size, s7.Buffer); err != nil {
		log.Panic("DB", s7.DbNumber, "couldn't be read", "\n", err)
	}

	return s7.Buffer

}

// ReadBoolByte exported function to be used to match boolean PLC data to DB schema.
func ReadBoolByte(s7 *S7, machineStates *controllerDataBase.MachineStates) *controllerDataBase.MachineStates {
	bufferSlice := s7.clientReadDb()
	bufferByte := bufferSlice[s7.Start]
	helper := gos7.Helper{}
	var indexBit uint
	fmt.Println("***************MachineStates*******************")
	for indexBit = 0; indexBit <= 5; indexBit++ {
		state := helper.GetBoolAt(bufferByte, indexBit)
		switch indexBit {
		case 0:
			machineStates.Running = state
			fmt.Println("Running = ", machineStates.Running)
		case 1:
			machineStates.AlertStop = state
			fmt.Println("AlertStop = ", machineStates.AlertStop)
		case 2:
			machineStates.FuncionalStop = state
			fmt.Println("FuncionalStop = ", machineStates.FuncionalStop)
		case 3:
			machineStates.StopByFoward = state
			fmt.Println("StopByFoward = ", machineStates.StopByFoward)
		case 4:
			machineStates.StopByBackward = state
			fmt.Println("StopByBackward = ", machineStates.StopByBackward)
		case 5:
			machineStates.UnderSpeed = state
			fmt.Println("UnderSpeed = ", machineStates.UnderSpeed)
		}

	}
	return machineStates
}
