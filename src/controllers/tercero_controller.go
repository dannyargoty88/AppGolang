package controllers

import (
	res "app/src/libs/response"
	models "app/src/models"
	interactor "app/src/usecase/interactor"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

//Builder
type terceroController struct{}

func NewTerceroController() TerceroController {
	return &terceroController{}
}

type TerceroController interface {
	GetAll(c echo.Context) error
	GetOne(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	SendEmail(c echo.Context) error
}

//Methods
func (ctrl terceroController) GetAll(c echo.Context) error {

	//get request data
	var params models.TbTerceroQuery
	if err := c.Bind(&params); err != nil {
		return res.RespError(c, err)
	}
	//validate request data
	if _, err_data := govalidator.ValidateStruct(params); err_data != nil {
		return res.RespError(c, err_data)
	}
	//process
	res_data, err := interactor.NewTerceroInteractor().GetAll(params)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Lista de Terceros", res_data)
}

func (ctrl terceroController) GetOne(c echo.Context) error {

	//get request data
	id := c.Param("id")
	//process
	res_data, err := interactor.NewTerceroInteractor().GetOne(id)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Información del Tercero", res_data)
}

func (ctrl terceroController) Create(c echo.Context) error {

	//get request data
	var params *models.TbTercero
	if err := c.Bind(&params); err != nil {
		return res.RespError(c, err)
	}
	//validate request data
	if _, err_data := govalidator.ValidateStruct(params); err_data != nil {
		return res.RespError(c, err_data)
	}
	//process
	res_data, err := interactor.NewTerceroInteractor().Create(params)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Tercero creado con éxito", res_data)
}

func (ctrl terceroController) Update(c echo.Context) error {

	//get request data
	var params *models.TbTercero
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
	res_data, err := interactor.NewTerceroInteractor().Update(id, params)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Tercero actualizado con éxito", res_data)
}

func (ctrl terceroController) Delete(c echo.Context) error {

	//get request data
	id := c.Param("id")
	//process
	res_data, err := interactor.NewTerceroInteractor().Delete(id)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Tercero borrado con éxito", res_data)
}

func (ctrl terceroController) SendEmail(c echo.Context) error {

	//get request data
	id := c.Param("id")

	//process
	res_data, err := interactor.NewTerceroInteractor().SendEmail(id)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Envio de correo de prueba", res_data)
}
