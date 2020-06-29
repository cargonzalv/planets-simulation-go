package model

import (
	"math"
)

/*
Planeta Estructura de un planeta
	nombre: Nombre del planeta
	velocidad: su velocidad en grados por día (velocidad indica movimiento horario,
	negativa indica movimiento anti-horario
	posicion: Coordenadas polares de la posición actual del planeta en cuestión
*/
type Planeta struct {
	nombre    string
	velocidad float64
	posicion  CoordsPolares
}

// RespuestaClimaGeneral Estructura de respuesta de consulta de clima general (para los ultimos 10 años)
type RespuestaClimaGeneral struct {
	Sequias        int `json:"sequias"`
	Lluvias        int `json:"lluvias"`
	DiaPicoLluvias int `json:"diaPicoLluvias"`
	Optimos        int `json:"optimos"`
}

// RespuestaClima Estructura de respuesta de consulta del clima del universo en un día particular
type RespuestaClima struct {
	Dia   int    `json:"dia"`
	Clima string `json:"clima"`
}

// CoordsPolares Estructura de las coordenadas polares de un planeta
type CoordsPolares struct {
	radio  float64
	angulo float64
}

func CrearPlaneta(nombre string, velocidad float64, radio float64) {
	p := Planeta{
		nombre:    nombre,
		velocidad: velocidad,
		posicion:  CoordsPolares{radio, 0},
	}
	agregarPlaneta(p)
}

func (p Planeta) CalcularDiasPorAnio() int {
	return int(math.Abs(360 / p.velocidad))
}
