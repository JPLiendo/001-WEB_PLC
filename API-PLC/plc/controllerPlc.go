package controllerPlc

import (
	"PLC-WEB/API-PLC/dataBase/modelDataBase"
	"log"
	"time"

	"github.com/robinson/gos7"
)

//Tipo S7Connection, parametriza la conexión con el PLC.
type S7Connection struct {
	Addr        string
	Rack        int
	Slot        int
	TimeOut     time.Duration
	IdleTimeout time.Duration
}

//Tipo S7DataRead, parametriza el Db a ser leido.
type S7DataRead struct {
	DbNumber int
	Start    int
	Size     int
	Buffer   []byte
}

// Tipo S7DataWrite struct {
// 	DbNumber int
// 	Start    int
// 	Size     int
// 	Buffer   []byte
// }

//Tipo S7Data, agrupa los datos leidos y escritos en el PLC.
type S7Data struct {
	S7DataRead
	// S7DataWrite
}

//Tipo S7, agrupa los datos de conexión y de lectura/escritura.
type S7 struct {
	S7Connection
	S7Data
}

//Interface, colección  de metodos.
type S7Connction interface {
	ClientReadDb()
}

//handlerConnection crea el handler para conectarse con el PLC.
func handlerConnection(s7 *S7Connection) *gos7.TCPClientHandler {
	handlerConnection := gos7.NewTCPClientHandler(s7.Addr, s7.Rack, s7.Slot)
	handlerConnection.Timeout = s7.TimeOut
	handlerConnection.IdleTimeout = s7.IdleTimeout
	//handlerConnection.Logger = log.New(os.Stdout, "tcp: ", log.LstdFlags)
	return handlerConnection
}

/*ClientReadDb lee un Db a partir de los parametros configurados en el tipo S7DataRead*/
func (s7 *S7) ClientReadDb() *modelDataBase.SensoresAnalogicos {
	//Inicio la conexión con el PLC y difiero el cierre.
	handler := handlerConnection(&s7.S7Connection)
	err := handler.Connect()
	if err != nil {
		log.Panic("No se pudo conectar con el PLC direccion: ", handler.Address, "\n", err)
	}
	defer handler.Close()

	//Creo el cliente y leo el Db.
	client := gos7.NewClient(handler)
	err = client.AGReadDB(s7.DbNumber, s7.Start, s7.Size, s7.Buffer)
	if err != nil {
		log.Panic("No se pudo leer el Db: ", s7.DbNumber, "\n", err)
	}
	//convierto los bytes a real y los guardo en el slice de sensores.
	s7Helper := gos7.Helper{}

	sensoresAnalogicos := modelDataBase.SensoresAnalogicos{}

	indexBuffer := 0
	for i := 0; i <= (len(s7.Buffer)/4)-1; i++ {
		sensorData := s7Helper.GetRealAt(s7.Buffer, indexBuffer)
		switch i {
		case 0:
			sensoresAnalogicos.SensorAnalogico_1 = sensorData
		case 1:
			sensoresAnalogicos.SensorAnalogico_2 = sensorData
		case 2:
			sensoresAnalogicos.SensorAnalogico_3 = sensorData
		case 3:
			sensoresAnalogicos.SensorAnalogico_4 = sensorData
		case 4:
			sensoresAnalogicos.SensorAnalogico_5 = sensorData
		case 5:
			sensoresAnalogicos.SensorAnalogico_6 = sensorData
		case 6:
			sensoresAnalogicos.SensorAnalogico_7 = sensorData
		case 7:
			sensoresAnalogicos.SensorAnalogico_8 = sensorData
		}

		indexBuffer = indexBuffer + 4
	}

	return &sensoresAnalogicos
}
