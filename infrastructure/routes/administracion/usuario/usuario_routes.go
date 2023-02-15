package usuario

import (
	ctrl "app/src/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	//Crud
	g.GET("/usuario", ctrl.NewUsuarioController().GetAll)
	g.GET("/usuario/:id", ctrl.NewUsuarioController().GetOne)
	g.POST("/usuario", ctrl.NewUsuarioController().Create)
	g.PUT("/usuario/:id", ctrl.NewUsuarioController().Update)
	g.DELETE("/usuario/:id", ctrl.NewUsuarioController().Delete)
}
