package departamento

import (
	ctrl "app/src/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	//Crud
	g.GET("/departamento", ctrl.NewDepartamentoController().GetAll)
	g.GET("/departamento/:id", ctrl.NewDepartamentoController().GetOne)
	g.POST("/departamento", ctrl.NewDepartamentoController().Create)
	g.PUT("/departamento/:id", ctrl.NewDepartamentoController().Update)
	g.DELETE("/departamento/:id", ctrl.NewDepartamentoController().Delete)

	//Custom
	g.GET("/departamento/ciudad", ctrl.NewDepartamentoController().GetAllDepartamentoCiudad)
}
