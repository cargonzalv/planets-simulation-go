package main

import "strings"

// Planetas Listado de planetas del sistema solar
var Planetas []Planeta

func agregarPlaneta(p Planeta) {
	Planetas = append(Planetas, p)
}

func buscarPlanetaPorNombre(nombre string) *Planeta {
	for _, value := range Planetas {
		if strings.EqualFold(value.nombre, nombre) {
			return &value
		}
	}
	return nil
}
