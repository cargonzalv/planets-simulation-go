package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"prueba-meli/route"
)

// estados
const (
	Sequia = "sequia"
	Lluvia = "lluvia"
	Optimo = "optimo"
	Normal = "normal"
)

func main() {

	crearPlaneta("Ferengi", -1, 500)
	crearPlaneta("Betasoide", -3, 2000)
	crearPlaneta("Vulcano", 5, 1000)

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	router := route.Init()
	router.Run(fasthttp.New(":8888"))
}

func simularAlDia(dia int) Triangulo {
	diaConvertido := float64(dia)
	for index, value := range Planetas {
		Planetas[index].posicion.angulo = math.Mod(value.posicion.angulo+(value.velocidad*diaConvertido), 360)
	}
	return Triangulo{Planetas[0].posicion, Planetas[1].posicion, Planetas[2].posicion}
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

func calcularClimaDia(dia int) RespuestaClima {
	triangulo := simularAlDia(dia)
	clima := darEstado(triangulo)
	return RespuestaClima{
		Clima: clima,
		Dia:   dia,
	}
}

func simulacion(dias int, p Planeta) RespuestaClimaGeneral {
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
		for index, value := range Planetas {
			Planetas[index].posicion.angulo = math.Mod(value.posicion.angulo+value.velocidad, 360)
		}
		estadoPrevio = estado
	}
	return RespuestaClimaGeneral{
		Sequias:        countSequia,
		Lluvias:        countLluvias,
		DiaPicoLluvias: diaPicoLluvias,
		Optimos:        countOptimo,
	}
}
