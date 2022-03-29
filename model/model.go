package model

type Dato struct {
	Id     int     `gorm:"primary_key"`
	Entero int     `json:"entero"`
	Real   float32 `json:"real"`
	Cadena string  `json:"cadena"`
}

var Datos []Dato
