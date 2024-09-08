package routes

import (
	"code-structure/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, pc *controllers.PackageController) {
	e.GET("/api/v1/packages", pc.GetAllPackages)
	e.GET("/api/v1/packages/:id", pc.GetPackageByID)
	e.POST("/api/v1/packages", pc.CreatePackage)
	e.PUT("/api/v1/packages/:id", pc.UpdatePackage)
	e.DELETE("/api/v1/packages/:id", pc.DeletePackage)
}
