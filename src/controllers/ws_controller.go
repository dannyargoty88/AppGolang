package controllers

import (
	res "app/src/libs/response"
	models "app/src/models"
	interactor "app/src/usecase/interactor"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

//Builder
type wsController struct{}

func NewWsController() WsController {
	return &wsController{}
}

type WsController interface {
	GetAll(c echo.Context) error
	GetOne(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

//Methods
func (ctrl wsController) GetAll(c echo.Context) error {

	//get request data
	var params models.TbWsQuery
	if err := c.Bind(&params); err != nil {
		return res.RespError(c, err)
	}
	//validate request data
	if _, err_data := govalidator.ValidateStruct(params); err_data != nil {
		return res.RespError(c, err_data)
	}
	//process
	res_data, err := interactor.NewWsInteractor().GetAll(params)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Lista Ws", res_data)
}

func (ctrl wsController) GetOne(c echo.Context) error {

	//get request data
	id := c.Param("id")
	//process
	res_data, err := interactor.NewWsInteractor().GetOne(id)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Información Ws", res_data)
}

func (ctrl wsController) Create(c echo.Context) error {

	//get request data
	var params *models.TbWs
	if err := c.Bind(&params); err != nil {
		return res.RespError(c, err)
	}
	//validate request data
	if _, err_data := govalidator.ValidateStruct(params); err_data != nil {
		return res.RespError(c, err_data)
	}
	//process
	res_data, err := interactor.NewWsInteractor().Create(params)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Ws creado con éxito", res_data)
}

func (ctrl wsController) Update(c echo.Context) error {

	//get request data
	var params *models.TbWs
	if err := c.Bind(&params); err != nil {
		return res.RespError(c, err)
	}
	//validate request data
	if _, err_data := govalidator.ValidateStruct(params); err_data != nil {
		return res.RespError(c, err_data)
	}
	//get request data
	id := c.Param("id")
	//process
	res_data, err := interactor.NewWsInteractor().Update(id, params)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Ws actualizado con éxito", res_data)
}

func (ctrl wsController) Delete(c echo.Context) error {

	//get request data
	id := c.Param("id")
	//process
	res_data, err := interactor.NewWsInteractor().Delete(id)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Ws borrado con éxito", res_data)
}
