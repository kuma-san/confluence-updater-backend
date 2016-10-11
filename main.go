package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/kuma-san/confluence-updater-backend/route"
	"github.com/labstack/echo/engine/fasthttp"
)

func init() {

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {

	router := route.Init()
	router.Run(fasthttp.New(":8888"))
}
