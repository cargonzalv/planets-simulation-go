package main

import (
	"log"
	"os"
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

	model.CrearPlanetas()

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	router := route.New()
	router.Start(":" + port)
}
