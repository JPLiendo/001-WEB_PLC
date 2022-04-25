package connectionPlc

import (
	"time"

	"github.com/JPLiendo/001-WEB_PLC/API-PLC/plc/controllerPlc"
)

var PLC = controllerPlc.S7{
	S7Connection: controllerPlc.S7Connection{
		Addr:        "192.168.0.10",
		Rack:        0,
		Slot:        2,
		TimeOut:     200 * time.Second,
		IdleTimeout: 200 * time.Second,
	},
	S7Data: controllerPlc.S7Data{
		S7DataRead: controllerPlc.S7DataRead{
			DbNumber: 102,
			Start:    0,
			Size:     1,
			Buffer:   make([]byte, 1),
		},
	},
}
