package modelDataBase

import "time"

//Modelo
type SensoresAnologicos struct {
	Id          int       `gorm:"primary_key"`
	ValorSensor []float64 `json:"valor_sensor"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
