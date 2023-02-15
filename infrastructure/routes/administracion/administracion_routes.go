package administracion_routes

import (
	ciudad "app/infrastructure/routes/administracion/ciudad"
	departamento "app/infrastructure/routes/administracion/departamento"
	tercero "app/infrastructure/routes/administracion/tercero"
	usuario "app/infrastructure/routes/administracion/usuario"

	"github.com/labstack/echo/v4"
)

//Routes private
func RoutesPrivate(private *echo.Group) {
	ciudad.Routes(private)
	departamento.Routes(private)
	tercero.Routes(private)
	usuario.Routes(private)
}
