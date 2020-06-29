package model

import (
	"fmt"
	"math"
	"strings"
)

// Planetas Listado de planetas del sistema solar
var Planetas []Planeta

func CrearPlanetas() {
	CrearPlaneta("Ferengi", -1, 500)
	CrearPlaneta("Betasoide", -3, 2000)
	CrearPlaneta("Vulcano", 5, 1000)
}

func agregarPlaneta(p Planeta) {
	Planetas = append(Planetas, p)
}

func BuscarPlanetaPorNombre(nombre string) *Planeta {
	for _, value := range Planetas {
		if strings.EqualFold(value.nombre, nombre) {
			return &value
		}
	}
	return nil
}

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

func Simulacion(dias int, p Planeta) RespuestaClimaGeneral {
	fmt.Println("Calculando simulación de ", dias, "días) para el Planeta", p.nombre)
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
			perimetro := darPerimetro(Planetas[0].posicion, Planetas[1].posicion, Planetas[2].posicion)
			maxPerimetro = math.Max(maxPerimetro, perimetro)
			if maxPerimetro == perimetro {
				diaPicoLluvias = i
			}
		} else if estado == Optimo {
			if estadoPrevio != estado {
				countOptimo++
			}
		}
		if i%100 == 0 {
			fmt.Println("dia:", i, estado, countSequia, countLluvias, countOptimo, maxPerimetro, diaPicoLluvias)
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
