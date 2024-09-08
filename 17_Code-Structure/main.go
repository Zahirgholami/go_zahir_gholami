package main

import (
	"code-structure/controllers"
	"code-structure/models"
	"code-structure/routes"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := gorm.Open("sqlite3", "packages.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	models.Migrate(db)

	e := echo.New()

	pc := &controllers.PackageController{DB: db}

	routes.RegisterRoutes(e, pc)
	e.Logger.Fatal(e.Start(":8080"))
}
