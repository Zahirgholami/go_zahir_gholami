package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Food struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

var foods = make([]Food, 0)

func main() {
	e := echo.New()

	e.GET("/api/v1/foods", getAllFoods)
	e.GET("/api/v1/foods/:id", getFoodByID)
	e.POST("/api/v1/foods", addFood)
	e.PUT("/api/v1/foods/:id", updateFood)
	e.DELETE("/api/v1/foods/:id", deleteFood)

	e.Logger.Fatal(e.Start(":8080"))
}

func getAllFoods(c echo.Context) error {
	return c.JSON(http.StatusOK, foods)
}

func getFoodByID(c echo.Context) error {
	id := c.Param("id")
	for _, food := range foods {
		if food.ID == id {
			return c.JSON(http.StatusOK, food)
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"message": "Food not found"})
}

func addFood(c echo.Context) error {
	food := new(Food)
	if err := c.Bind(food); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid input"})
	}
	food.ID = uuid.New().String()
	foods = append(foods, *food)
	return c.JSON(http.StatusCreated, food)
}

func updateFood(c echo.Context) error {
	id := c.Param("id")
	for i, food := range foods {
		if food.ID == id {
			if err := c.Bind(&food); err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid input"})
			}
			foods[i] = food
			return c.JSON(http.StatusOK, food)
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"message": "Food not found"})
}

func deleteFood(c echo.Context) error {
	id := c.Param("id")
	for i, food := range foods {
		if food.ID == id {
			foods = append(foods[:i], foods[i+1:]...)
			return c.JSON(http.StatusOK, echo.Map{"message": "Food deleted"})
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"message": "Food not found"})
}
