package route

import (
	"prueba-meli/handler"

	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {

	e := echo.New()

	// Set Bundle MiddleWare
	e.Use(echoMw.Logger())
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	e.HTTPErrorHandler = handler.JSONHTTPErrorHandler

	// Routes
	v1 := e.Group("/")
	{
		v1.GET("/", api.indexHandler())
		v1.GET("/clima", api.climaHandler())
		v1.GET("/simulacion", api.simulacionHandler())
	}
	return e
}
