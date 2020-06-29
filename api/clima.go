package api

import (
	"fmt"
	"net/http"
	"prueba-meli/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

// ClimaHandler Responde a nuestro request con la información del clima para el día especificado
// @Summary Devuelve la información del clima para el día especificado
// @Produce json
// @Param dia query int true "Dia"
// @Success 200 {object} model.RespuestaClima
// @Failure 400 {object} model.HTTPError
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /clima [get]
func ClimaHandler(c echo.Context) error {
	if c.QueryParam("dia") == "" {
		return echo.NewHTTPError(400, "Debe especificar el día a buscar")
	}
	dia, error := strconv.Atoi(c.QueryParam("dia"))
	if error != nil {
		return echo.NewHTTPError(400, "Debe especificar el día como un número")
	}
	clima := model.BuscarClimaDia(dia)
	fmt.Printf("%+v\n", clima)
	return c.JSON(http.StatusOK, clima)
}
