package controllers

import (
	res "app/src/libs/response"
	models "app/src/models"
	interactor "app/src/usecase/interactor"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

//Builder
type usuarioController struct{}

func NewUsuarioController() UsuarioController {
	return &usuarioController{}
}

type UsuarioController interface {
	GetAll(c echo.Context) error
	GetOne(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

//Methods
func (ctrl usuarioController) GetAll(c echo.Context) error {

	//get request data
	var params models.TbUsuarioQuery
	if err := c.Bind(&params); err != nil {
		return res.RespError(c, err)
	}
	//validate request data
	if _, err_data := govalidator.ValidateStruct(params); err_data != nil {
		return res.RespError(c, err_data)
	}
	//process
	res_data, err := interactor.NewUsuarioInteractor().GetAll(params)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Lista de usuarios", res_data)
}

func (ctrl usuarioController) GetOne(c echo.Context) error {

	//get request data
	id := c.Param("id")
	//process
	res_data, err := interactor.NewUsuarioInteractor().GetOne(id)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Información del usuario", res_data)
}

func (ctrl usuarioController) Create(c echo.Context) error {

	//get request data
	var params *models.TbUsuario
	if err := c.Bind(&params); err != nil {
		return res.RespError(c, err)
	}
	//validate request data
	if _, err_data := govalidator.ValidateStruct(params); err_data != nil {
		return res.RespError(c, err_data)
	}
	//process
	res_data, err := interactor.NewUsuarioInteractor().Create(params)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Usuario creado con éxito", res_data)
}

func (ctrl usuarioController) Update(c echo.Context) error {

	//get request data
	var params *models.TbUsuario
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
	res_data, err := interactor.NewUsuarioInteractor().Update(id, params)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Usuario actualizado con éxito", res_data)
}

func (ctrl usuarioController) Delete(c echo.Context) error {

	//get request data
	id := c.Param("id")
	//process
	res_data, err := interactor.NewUsuarioInteractor().Delete(id)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Usuario borrado con éxito", res_data)
}
