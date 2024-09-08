package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddleware() echo.MiddlewareFunc {
	return middleware.Logger()
}
