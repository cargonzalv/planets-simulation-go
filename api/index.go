package api

import (
	"github.com/labstack/echo"
)

// indexHandler Responde a nuestro request con un saludo
func indexHandler(c echo.Context) error {
	return c.String(200, "Hello, Universe!")
}
