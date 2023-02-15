package index

import (
	ctrl "app/src/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	g.GET("/", ctrl.NewIndexController().Index)
}
