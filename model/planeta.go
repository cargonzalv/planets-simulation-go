package model

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
