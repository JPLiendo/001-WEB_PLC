package modelDataBase

import "time"

type SensorAnalogico struct {
	Id        int       `gorm:"primary_key"`
	Value     float32   `json:"value"`
	UpdatedAt time.Time `json:"updatedAt"`
}
