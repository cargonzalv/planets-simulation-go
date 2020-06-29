package route

import (
	"prueba-meli/api"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

// New Inicialización del servidor
func New() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	// Set Bundle MiddleWare
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Validator = NewValidator()

	// Routes
	e.GET("/", api.IndexHandler)
	e.GET("/clima", api.ClimaHandler)
	e.GET("/simulacion", api.SimulacionHandler)
	return e
}
