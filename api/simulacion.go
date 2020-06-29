package api

import (
	"fmt"
	"net/http"
	"prueba-meli/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

// SimulacionHandler Responde a nuestro request con la información de la simulación
// @Summary Devuelve la información de una simulación de 10 años (por defecto, se puede cambiar)
// @Produce json
// @Param anios query int false "Años a simular"
// @Param planeta query string false "Planeta en el que se basan los años"
// @Success 200 {object} model.RespuestaClimaGeneral
// @Failure 400 {object} model.HTTPError
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /api/simulacion [get]
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
