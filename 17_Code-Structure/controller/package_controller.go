package controllers

import (
	"net/http"
	"strconv"

	"code-structure/models"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type PackageController struct {
	DB *gorm.DB
}

func (pc *PackageController) GetAllPackages(c echo.Context) error {
	var packages []models.Package
	pc.DB.Find(&packages)
	return c.JSON(http.StatusOK, packages)
}

func (pc *PackageController) GetPackageByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var pkg models.Package
	if err := pc.DB.First(&pkg, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Package not found"})
	}
	return c.JSON(http.StatusOK, pkg)
}

func (pc *PackageController) CreatePackage(c echo.Context) error {
	var pkg models.Package
	if err := c.Bind(&pkg); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	pc.DB.Create(&pkg)
	return c.JSON(http.StatusCreated, pkg)
}

func (pc *PackageController) UpdatePackage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var pkg models.Package
	if err := pc.DB.First(&pkg, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Package not found"})
	}

	if err := c.Bind(&pkg); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	pc.DB.Save(&pkg)
	return c.JSON(http.StatusOK, pkg)
}

func (pc *PackageController) DeletePackage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := pc.DB.Delete(&models.Package{}, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Package not found"})
	}
	return c.NoContent(http.StatusNoContent)
}
