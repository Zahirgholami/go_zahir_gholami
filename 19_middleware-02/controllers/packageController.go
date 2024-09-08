package controllers

import (
	"middleware-02/config"
	"middleware-02/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPackages(c echo.Context) error {
	packages := []models.Package{}
	config.DB.Find(&packages)
	return c.JSON(http.StatusOK, packages)
}

func GetPackageByID(c echo.Context) error {
	id := c.Param("id")
	packageItem := new(models.Package)
	config.DB.First(&packageItem, id)
	if packageItem.ID == 0 {
		return c.JSON(http.StatusNotFound, "Package not found")
	}
	return c.JSON(http.StatusOK, packageItem)
}

func CreatePackage(c echo.Context) error {
	packageItem := new(models.Package)
	if err := c.Bind(packageItem); err != nil {
		return err
	}

	config.DB.Create(&packageItem)
	return c.JSON(http.StatusCreated, packageItem)
}

func UpdatePackage(c echo.Context) error {
	id := c.Param("id")
	packageItem := new(models.Package)
	config.DB.First(&packageItem, id)
	if packageItem.ID == 0 {
		return c.JSON(http.StatusNotFound, "Package not found")
	}

	if err := c.Bind(packageItem); err != nil {
		return err
	}

	config.DB.Save(&packageItem)
	return c.JSON(http.StatusOK, packageItem)
}

func DeletePackage(c echo.Context) error {
	id := c.Param("id")
	packageItem := new(models.Package)
	config.DB.First(&packageItem, id)
	if packageItem.ID == 0 {
		return c.JSON(http.StatusNotFound, "Package not found")
	}

	config.DB.Delete(&packageItem)
	return c.JSON(http.StatusNoContent, nil)
}
