package controllers

import (
	res "app/src/libs/response"
	models "app/src/models"
	interactor "app/src/usecase/interactor"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

//Builder
type itemWsController struct{}

func NewItemWsController() ItemWsController {
	return &itemWsController{}
}

type ItemWsController interface {
	GetAll(c echo.Context) error
	GetOne(c echo.Context) error
}

//Methods
func (ctrl itemWsController) GetAll(c echo.Context) error {

	//get request data
	var params models.TbItemwsQuery
	if err := c.Bind(&params); err != nil {
		return res.RespError(c, err)
	}
	//validate request data
	if _, err_data := govalidator.ValidateStruct(params); err_data != nil {
		return res.RespError(c, err_data)
	}
	//process
	res_data, err := interactor.NewItemWsInteractor().GetAll(params)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Lista de Items Ws", res_data)
}

func (ctrl itemWsController) GetOne(c echo.Context) error {

	//get request data
	ws := c.Param("ws")
	id := c.Param("id")
	//process
	res_data, err := interactor.NewItemWsInteractor().GetOne(ws, id)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Información del Items Ws", res_data)
}
