package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// climaHandler Responde a nuestro request con la información del clima
func climaHandler(c echo.Context) error {
	if c.QueryParam("dia") == "" {
		return echo.NewHTTPError(400, "Debe especificar el día a buscar")
	}
	dia, error := strconv.Atoi(c.QueryParam("dia"))
	if error != nil {
		return echo.NewHTTPError(400, "Debe especificar el día como un número")
	}
	clima := calcularClimaDia(dia)
	fmt.Printf("%+v\n", clima)
	return c.JSON(http.StatusOK, clima)
}
