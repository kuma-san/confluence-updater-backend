package route

import (
	"github.com/kuma-san/confluence-updater-backend/api"
	"github.com/kuma-san/confluence-updater-backend/handler"
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Debug()

	// Set Bundle MiddleWare
	e.Use(echoMw.Logger())
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.GET("/status", api.GetStatus())
		v1.POST("/task", api.PostTask())

	}
	return e
}
