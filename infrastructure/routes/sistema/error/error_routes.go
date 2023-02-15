package error

import (
	ctrl "app/src/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	//Crud
	g.GET("/error", ctrl.NewErrorController().GetAll)
	g.GET("/error/:id", ctrl.NewErrorController().GetOne)
	g.DELETE("/error/:id", ctrl.NewErrorController().Delete)
}
