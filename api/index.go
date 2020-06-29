package api

import (
	"github.com/labstack/echo"
)

// IndexHandler Responde a nuestro request con un saludo
func IndexHandler(c echo.Context) error {
	return c.String(200, "Hello, Universe!")
}
