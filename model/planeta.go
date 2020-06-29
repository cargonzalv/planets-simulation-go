package model

import "math"

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

/*
PlanetaDB Estructura de un planeta en la DB
	nombre: Nombre del planeta
	velocidad: su velocidad en grados por día (velocidad indica movimiento horario,
	negativa indica movimiento anti-horario
*/
type PlanetaDB struct {
	Nombre    string
	Velocidad float64
	Radio     float64
}

// CalcularDiasPorAnio Calcula los días que contiene un año para el planeta receiver
func (p Planeta) CalcularDiasPorAnio() int {
	return int(math.Abs(360 / p.velocidad))
}
