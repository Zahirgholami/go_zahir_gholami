package routes

import (
	"middleware-02/controllers"
	"middleware-02/middlewares"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/api/v1/register", controllers.Register)
	e.POST("/api/v1/login", controllers.Login)

	r := e.Group("/api/v1/packages")
	r.Use(middlewares.JWTMiddleware())

	r.GET("", controllers.GetPackages)
	r.GET("/:id", controllers.GetPackageByID)
	r.POST("", controllers.CreatePackage)
	r.PUT("/:id", controllers.UpdatePackage)
	r.DELETE("/:id", controllers.DeletePackage)
}
