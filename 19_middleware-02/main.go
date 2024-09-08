package main

import (
	"middleware-02/config"
	"middleware-02/middlewares"
	"middleware-02/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.InitDB()

	e.Use(middlewares.LogMiddleware())

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
