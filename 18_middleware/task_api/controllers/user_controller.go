package controllers

import (
	"net/http"
	"task_api/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	db.Create(&user)
	return c.JSON(http.StatusOK, user)
}

func Login(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	dbUser := new(models.User)
	db.Where("email = ?", user.Email).First(&dbUser)

	if dbUser.ID == 0 {
		return c.JSON(http.StatusUnauthorized, "invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, "invalid credentials")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = dbUser.ID

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
