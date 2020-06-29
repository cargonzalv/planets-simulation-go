package model

import (
	"fmt"
	"math"
	"strings"
)

// Variables Globales

// Planetas Listado de planetas del sistema solar
var Planetas []Planeta
var cacheClima map[int]RespuestaClimaGeneral = make(map[int]RespuestaClimaGeneral)

// CrearPlaneta Crea un nuevo planeta con su nombre, velocidad y radio ingresados por parámetro
func CrearPlaneta(nombre string, velocidad float64, radio float64) Planeta {
	p := Planeta{
		nombre:    nombre,
		velocidad: velocidad,
		posicion:  CoordsPolares{radio, 0},
	}
	agregarPlaneta(p)
	return p
}

// CrearPlanetas Crea los planetas del sistema solar descrito en el enunciado
func CrearPlanetas() {
	CrearPlaneta("Ferengi", -1, 500)
	CrearPlaneta("Betasoide", -3, 2000)
	CrearPlaneta("Vulcano", 5, 1000)
}

func agregarPlaneta(p Planeta) {
	Planetas = append(Planetas, p)
}

// BuscarPlanetaPorNombre Busca el struct Planeta por su nombre
func BuscarPlanetaPorNombre(nombre string) *Planeta {
	for _, value := range Planetas {
		if strings.EqualFold(value.nombre, nombre) {
			return &value
		}
	}
	return nil
}

// CalcularClimaDia Calcula el estado del clima para el día ingresado por parámetro
func CalcularClimaDia(dia int) RespuestaClima {
	triangulo := avanzarDias(dia)
	fmt.Printf("%+v", triangulo)
	clima := darEstado(triangulo)
	return RespuestaClima{
		Clima: clima,
		Dia:   dia,
	}
}

func darEstado(t Triangulo) string {
	if t.sonColineales() {
		if t.sonColinealesCentro() {
			return Sequia
		}
		return Optimo
	} else if t.trianguloContieneOrigen() {
		return Lluvia
	} else {
		return Normal
	}
}

// Simulacion hace el proceso de avanzar en el tiempo n días y calcular el conteo de clima
func Simulacion(dias int) RespuestaClimaGeneral {
	fmt.Println("Calculando simulación de ", dias, "días")
	if cacheClima[dias] != (RespuestaClimaGeneral{}) {
		return cacheClima[dias]
	}
	countSequia := 0
	countLluvias := 0
	countOptimo := 0
	diaPicoLluvias := 0
	estadoPrevio := ""
	var maxPerimetro float64 = -1
	for i := 1; i <= dias; i++ {
		triangulo := Triangulo{Planetas[0].posicion, Planetas[1].posicion, Planetas[2].posicion}
		estado := darEstado(triangulo)
		if estado == Sequia {
			if estadoPrevio != estado {
				countSequia++
			}
		} else if estado == Lluvia {
			if estadoPrevio != estado {
				countLluvias++
			}
			perimetro := DarPerimetro(Planetas[0].posicion, Planetas[1].posicion, Planetas[2].posicion)
			maxPerimetro = math.Max(maxPerimetro, perimetro)
			if maxPerimetro == perimetro {
				diaPicoLluvias = i
			}
		} else if estado == Optimo {
			if estadoPrevio != estado {
				countOptimo++
			}
		}
		cacheClima[i] = RespuestaClimaGeneral{
			Sequias:        countSequia,
			Lluvias:        countLluvias,
			DiaPicoLluvias: diaPicoLluvias,
			Optimos:        countOptimo,
		}
		avanzarDias(1)
		estadoPrevio = estado
	}
	return RespuestaClimaGeneral{
		Sequias:        countSequia,
		Lluvias:        countLluvias,
		DiaPicoLluvias: diaPicoLluvias,
		Optimos:        countOptimo,
	}
}

func avanzarDias(dia int) Triangulo {
	for index, value := range Planetas {
		Planetas[index].posicion.angulo = math.Mod(value.posicion.angulo+(value.velocidad*float64(dia)), 360)
	}
	return Triangulo{Planetas[0].posicion, Planetas[1].posicion, Planetas[2].posicion}
}
