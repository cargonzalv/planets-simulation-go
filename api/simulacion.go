package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// simulacionHandler Responde a nuestro request con la información de la simulación
func simulacionHandler(c echo.Context) error {

	anios := 10            // Por defecto la simulación es de 10 años
	planeta := Planetas[0] // Por defecto la simulación es de los años de Ferengi (Planeta mas lento)
	if a, err := strconv.Atoi(c.QueryParam("anios")); err != nil {
		if err == nil {
			anios = a
		}
	}
	if nombrePlaneta := c.QueryParam("planeta"); nombrePlaneta != "" {
		if p := buscarPlanetaPorNombre(nombrePlaneta); p != nil {
			planeta = *p
		}
	}
	dias := anios * planeta.calcularDiasPorAnio()
	respuesta := simulacion(dias, planeta)
	fmt.Printf("%+v\n", respuesta)
	return c.JSON(http.StatusOK, respuesta)
}
