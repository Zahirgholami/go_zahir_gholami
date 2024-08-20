package routes

import (
	"task_api/controllers"
	"task_api/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/api/v1/register", controllers.Register)
	e.POST("/api/v1/login", controllers.Login)

	p := e.Group("/api/v1/packages")
	p.Use(middleware.JWTMiddleware())

	p.GET("", controllers.GetPackages)
	p.GET("/:id", controllers.GetPackageByID)
	p.POST("", controllers.CreatePackage)
	p.PUT("/:id", controllers.UpdatePackage)
	p.DELETE("/:id", controllers.DeletePackage)
}
