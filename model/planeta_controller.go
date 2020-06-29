package model

import (
	"fmt"
	"math"
	"prueba-meli/db"
	"strings"
)

// Variables Globales

// Planetas Listado de planetas del sistema solar
var Planetas []Planeta

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
	var err error
	planetas := []PlanetaDB{}
	err = db.Database.Table("planeta").Model(&PlanetaDB{}).Limit(100).Find(&planetas).Error
	fmt.Println(planetas, err)
	if err != nil || len(planetas) == 0 {
		panic("Error: No fue posible encontrar los planetas del sistema solar")
	}
	CrearPlaneta(planetas[0].Nombre, planetas[0].Velocidad, planetas[0].Radio)
	CrearPlaneta(planetas[1].Nombre, planetas[1].Velocidad, planetas[1].Radio)
	CrearPlaneta(planetas[2].Nombre, planetas[2].Velocidad, planetas[2].Radio)
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
func Simulacion(dias int, job bool) RegistroClima {
	var ultimoDia int = 0
	errorSelect := db.Database.Raw("SELECT dia FROM registroclima ORDER BY dia desc LIMIT 1", &ultimoDia).Error
	if errorSelect != nil {
		fmt.Println("Error en select de registros clima", errorSelect)
	}
	// Si ya hay al menos un día procesado, estamos en el job y el ultimoDia + 1 es mayor a los días
	// a calcular, aumentamos los días más uno
	if ultimoDia > 0 && job {
		dias = int(math.Max(float64(dias), float64(ultimoDia)+1))
	}
	fmt.Println("Calculando simulación de ", dias, "días")
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
		registro := RegistroClima{
			Sequias:        countSequia,
			Lluvias:        countLluvias,
			DiaPicoLluvias: diaPicoLluvias,
			Optimos:        countOptimo,
			Clima:          estadoPrevio,
			Dia:            i,
		}

		if i > ultimoDia && job {
			// Si el dia actual es mayor al ultimo dia en la db y estamos en el job, inserto este día
			errorInserting := db.Database.Table("registroclima").Create(registro).Error
			if errorInserting != nil {
				fmt.Printf("Error en: %+v", registro)
			}
		}

		avanzarDias(1)
		estadoPrevio = estado
	}
	return RegistroClima{
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
