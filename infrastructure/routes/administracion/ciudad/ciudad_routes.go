package ciudad

import (
	ctrl "app/src/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	//Crud
	g.GET("/ciudad", ctrl.NewCiudadController().GetAll)
	g.GET("/ciudad/:id", ctrl.NewCiudadController().GetOne)
	g.POST("/ciudad", ctrl.NewCiudadController().Create)
	g.PUT("/ciudad/:id", ctrl.NewCiudadController().Update)
	g.DELETE("/ciudad/:id", ctrl.NewCiudadController().Delete)
}
