package api

import (
	"fmt"
	"net/http"
	"prueba-meli/model"
	"strconv"

	"github.com/labstack/echo"
)

// SimulacionHandler Responde a nuestro request con la información de la simulación
func SimulacionHandler(c echo.Context) error {

	anios := 10                  // Por defecto la simulación es de 10 años
	planeta := model.Planetas[0] // Por defecto la simulación es de los años de Ferengi (Planeta mas lento)
	if a, err := strconv.Atoi(c.QueryParam("anios")); err == nil {
		anios = a
	}
	if nombrePlaneta := c.QueryParam("planeta"); nombrePlaneta != "" {
		if p := model.BuscarPlanetaPorNombre(nombrePlaneta); p != nil {
			planeta = *p
		}
	}
	dias := anios * planeta.CalcularDiasPorAnio()
	respuesta := model.Simulacion(dias, planeta)
	fmt.Printf("%+v\n", respuesta)
	return c.JSON(http.StatusOK, respuesta)
}
