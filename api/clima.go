package api

import (
	"fmt"
	"net/http"
	"prueba-meli/model"
	"strconv"

	"github.com/labstack/echo"
)

// ClimaHandler Responde a nuestro request con la información del clima
func ClimaHandler(c echo.Context) error {
	if c.QueryParam("dia") == "" {
		return echo.NewHTTPError(400, "Debe especificar el día a buscar")
	}
	dia, error := strconv.Atoi(c.QueryParam("dia"))
	if error != nil {
		return echo.NewHTTPError(400, "Debe especificar el día como un número")
	}
	clima := model.CalcularClimaDia(dia)
	fmt.Printf("%+v\n", clima)
	return c.JSON(http.StatusOK, clima)
}
