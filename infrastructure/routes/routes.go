package routes

import (
	middleware "app/infrastructure/middleware"
	administracion "app/infrastructure/routes/administracion"
	sistema "app/infrastructure/routes/sistema"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) *echo.Echo {

	//Routes public
	public := e.Group("/api")
	sistema.RoutesPublic(public)

	//Routes private
	private := e.Group("/api")
	private.Use(middleware.MiddlewareTokenApp())
	sistema.RoutesPrivate(private)
	administracion.RoutesPrivate(private)

	return e
}
