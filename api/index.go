package api

import (
	"github.com/labstack/echo/v4"
)

// IndexHandler Responde a nuestro request con un saludo
// @Summary Devuelve un saludo
// @Produce plain
// @Success 200 {string} string "Hello, Universe!"
// @Failure 500 {object} model.HTTPError
// @Router /api/ [get]
func IndexHandler(c echo.Context) error {
	return c.String(200, "Hello, Universe!")
}
