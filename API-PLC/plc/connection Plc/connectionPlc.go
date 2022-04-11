package connectionPlc

import (
	controllerPlc "PLC-WEB/API-PLC/plc/controller.plc"
	"time"
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
