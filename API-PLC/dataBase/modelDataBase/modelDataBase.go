package modelDataBase

import "time"

type SensoresAnalogicos struct {
	Id                int `gorm:"primary_key"`
	SensorAnalogico_1 float32
	SensorAnalogico_2 float32
	SensorAnalogico_3 float32
	SensorAnalogico_4 float32
	SensorAnalogico_5 float32
	SensorAnalogico_6 float32
	SensorAnalogico_7 float32
	SensorAnalogico_8 float32
	UpdatedAt         time.Time
}
