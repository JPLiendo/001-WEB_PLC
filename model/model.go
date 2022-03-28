package model

type Dato struct {
	Id     int     `json:"id,omitempty"`
	Entero int     `json:"entero"`
	Real   float32 `json:"real"`
	Cadena string  `json:"cadena"`
}

var Datos []Dato
