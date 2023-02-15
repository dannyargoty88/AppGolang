package tercero

import (
	ctrl "app/src/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	//Crud
	g.GET("/tercero", ctrl.NewTerceroController().GetAll)
	g.GET("/tercero/:id", ctrl.NewTerceroController().GetOne)
	g.POST("/tercero", ctrl.NewTerceroController().Create)
	g.PUT("/tercero/:id", ctrl.NewTerceroController().Update)
	g.DELETE("/tercero/:id", ctrl.NewTerceroController().Delete)
	g.PUT("/tercero/sendmail/:id", ctrl.NewTerceroController().SendEmail)
}
