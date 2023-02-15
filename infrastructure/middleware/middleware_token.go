package middleware_app

import (
	"app/src/libs/token"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MiddlewareTokenApp() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: token.SigningMethod().Name,
		SigningKey:    token.GetSigningKey(),
	})
}
