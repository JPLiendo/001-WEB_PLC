package model

import "time"

type SensoresAnalogicos struct {
	Id                int       `json:"id"`
	SensorAnalogico_1 float32   `json:"sensoranalogico_1"`
	SensorAnalogico_2 float32   `json:"sensoranalogico_2"`
	SensorAnalogico_3 float32   `json:"sensoranalogico_3"`
	SensorAnalogico_4 float32   `json:"sensoranalogico_4"`
	SensorAnalogico_5 float32   `json:"sensoranalogico_5"`
	SensorAnalogico_6 float32   `json:"sensoranalogico_6"`
	SensorAnalogico_7 float32   `json:"sensoranalogico_7"`
	SensorAnalogico_8 float32   `json:"sensoranalogico_8"`
	UpdatedAt         time.Time `json:"updateat"`
}

var DatosSensor []SensoresAnalogicos
