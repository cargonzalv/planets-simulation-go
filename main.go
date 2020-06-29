package main

import (
	"log"
	"os"
	"prueba-meli/db"
	"prueba-meli/model"
	"prueba-meli/route"
)

// @title Swagger Example API
// @version 1.0
// @description This is a weather calculator for planets.
// @termsOfService http://swagger.io/terms/

// @host https://ml-solar-system-281804.rj.r.appspot.com
// @BasePath /api

func main() {

	// Inicializamos conexion con base de datos
	db.Init()

	// Creamos los planetas del enunciado
	model.CrearPlanetas()

	// Configuración del puerto a escuchar
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// Inicialización de router
	router := route.New()
	// Empezamos a escuchar en el puerto seleccionado
	router.Start(":" + port)
}
