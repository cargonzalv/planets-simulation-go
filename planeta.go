package main

import (
	"math"
)

// Planeta bla
type Planeta struct {
	nombre    string
	velocidad float64
	posicion  CoordsPolares
}

type Respuesta struct {
	Sequias        int `json:"sequias"`
	Lluvias        int `json:"lluvias"`
	Optimos        int `json:"optimos"`
	DiaPicoLluvias int `json:"diaPicoLluvias"`
}

// CoordsPolares lba
type CoordsPolares struct {
	radio  float64
	angulo float64
}

func crearPlaneta(nombre string, velocidad float64, radio float64) {
	p := Planeta{
		nombre:    nombre,
		velocidad: velocidad,
		posicion:  CoordsPolares{radio, 0},
	}
	Planetas = append(Planetas, p)
}

func (p Planeta) calcularDiasPorAnio() int {
	return int(math.Abs(360 / p.velocidad))
}
