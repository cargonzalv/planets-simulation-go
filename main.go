package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
)

// estados
const (
	Sequia = iota
	Lluvia = iota
	Optimo = iota
	Normal = iota
)

// Planetas Listado de planetas del sistema solar
var Planetas []Planeta

func main() {

	crearPlaneta("Ferengi", 1, 500)
	crearPlaneta("Betasoide", 3, 2000)
	crearPlaneta("Vulcano", -5, 1000)

	http.HandleFunc("/", indexHandler)

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	countSequia, countLluvias, countOptimo, diaPicoLluvias := simulacion(10, Planetas[2])
	resp := Respuesta{
		Sequias:        countSequia,
		Lluvias:        countLluvias,
		Optimos:        countOptimo,
		DiaPicoLluvias: diaPicoLluvias,
	}
	fmt.Printf("%+v\n", resp)
	respJSON, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}

func simulacion(anios int, p Planeta) (int, int, int, int) {
	dias := anios * p.calcularDiasPorAnio()
	fmt.Println("Calculando simulación de", anios, "años (", dias, "días) para el Planeta", p.nombre)
	countSequia := 0
	countLluvias := 0
	countOptimo := 0
	diaPicoLluvias := 0
	estadoPrevio := 0
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
			countLluvias++
			perimetro := darPerimetro(Planetas[0].posicion, Planetas[1].posicion, Planetas[2].posicion)
			fmt.Println(perimetro)
			maxPerimetro = math.Max(maxPerimetro, perimetro)
			if maxPerimetro == perimetro {
				diaPicoLluvias = i
			}
		} else if estado == Optimo {
			if estadoPrevio != estado {
				countOptimo++
			}
		}
		fmt.Println("dia:", i, estado, countSequia, countLluvias, countOptimo, maxPerimetro, diaPicoLluvias)
		for index, value := range Planetas {
			Planetas[index].posicion.angulo = math.Mod(value.posicion.angulo+value.velocidad, 360)
		}
		estadoPrevio = estado
	}
	return countSequia, countLluvias, countOptimo, diaPicoLluvias
}

func darEstado(t Triangulo) int {
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
