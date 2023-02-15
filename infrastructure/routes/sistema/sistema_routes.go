package sistema_routes

import (
	createendpoint "app/infrastructure/routes/sistema/createendpoint"
	errores "app/infrastructure/routes/sistema/error"
	index "app/infrastructure/routes/sistema/index"
	login "app/infrastructure/routes/sistema/login"
	ws "app/infrastructure/routes/sistema/ws"

	"github.com/labstack/echo/v4"
)

//Routes public
func RoutesPublic(public *echo.Group) {
	index.Routes(public)
	login.Routes(public)
}

//Routes private
func RoutesPrivate(private *echo.Group) {
	createendpoint.Routes(private)
	errores.Routes(private)
	ws.Routes(private)
}
