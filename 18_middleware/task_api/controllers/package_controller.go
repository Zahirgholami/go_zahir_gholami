package controllers

import (
	"net/http"
	"task_api/models"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func GetPackages(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	var packages []models.Package
	db.Find(&packages)
	return c.JSON(http.StatusOK, packages)
}

func GetPackageByID(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	id := c.Param("id")
	var pkg models.Package
	if err := db.First(&pkg, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "package not found")
	}
	return c.JSON(http.StatusOK, pkg)
}

func CreatePackage(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	pkg := new(models.Package)
	if err := c.Bind(pkg); err != nil {
		return err
	}
	db.Create(&pkg)
	return c.JSON(http.StatusCreated, pkg)
}

func UpdatePackage(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	id := c.Param("id")
	var pkg models.Package
	if err := db.First(&pkg, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "package not found")
	}

	if err := c.Bind(&pkg); err != nil {
		return err
	}
	db.Save(&pkg)
	return c.JSON(http.StatusOK, pkg)
}

func DeletePackage(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	id := c.Param("id")
	if err := db.Delete(&models.Package{}, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "package not found")
	}
	return c.NoContent(http.StatusNoContent)
}
