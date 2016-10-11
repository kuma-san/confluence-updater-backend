package api

import (

	//"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

type ()

func GetStatus() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.NoContent(fasthttp.StatusOK)
	}
}
