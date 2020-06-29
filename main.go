package main

import (
	"log"
	"os"
	"prueba-meli/model"
	"prueba-meli/route"
)

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
