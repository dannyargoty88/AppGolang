package ws

import (
	ctrl "app/src/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {

	//Ws
	g.GET("/ws", ctrl.NewWsController().GetAll)
	g.GET("/ws/:id", ctrl.NewWsController().GetOne)
	g.POST("/ws", ctrl.NewWsController().Create)
	g.PUT("/ws/:id", ctrl.NewWsController().Update)
	g.DELETE("/ws/:id", ctrl.NewWsController().Delete)

	//Items
	g.GET("/itemws", ctrl.NewItemWsController().GetAll)
	g.GET("/itemws/:ws/:id", ctrl.NewItemWsController().GetOne)
}
