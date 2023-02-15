package controllers

import (
	res "app/src/libs/response"
	models "app/src/models"
	interactor "app/src/usecase/interactor"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

//Builder
type departamentoController struct{}

func NewDepartamentoController() DepartamentoController {
	return &departamentoController{}
}

type DepartamentoController interface {
	GetAll(c echo.Context) error
	GetOne(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	GetAllDepartamentoCiudad(c echo.Context) error
}

//Methods
func (ctrl departamentoController) GetAll(c echo.Context) error {

	//get request data
	var params models.TbDepartamentoQuery
	if err := c.Bind(&params); err != nil {
		return res.RespError(c, err)
	}
	//validate request data
	if _, err_data := govalidator.ValidateStruct(params); err_data != nil {
		return res.RespError(c, err_data)
	}
	//process
	res_data, err := interactor.NewDepartamentoInteractor().GetAll(params)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Lista de departamentos", res_data)
}

func (ctrl departamentoController) GetOne(c echo.Context) error {

	//get request data
	id := c.Param("id")
	//process
	res_data, err := interactor.NewDepartamentoInteractor().GetOne(id)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Información del departamento", res_data)
}

func (ctrl departamentoController) Create(c echo.Context) error {

	//get request data
	var params *models.TbDepartamento
	if err := c.Bind(&params); err != nil {
		return res.RespError(c, err)
	}
	//validate request data
	if _, err_data := govalidator.ValidateStruct(params); err_data != nil {
		return res.RespError(c, err_data)
	}
	//process
	res_data, err := interactor.NewDepartamentoInteractor().Create(params)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Departamento creado con éxito", res_data)
}

func (ctrl departamentoController) Update(c echo.Context) error {

	//get request data
	var params *models.TbDepartamento
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
	res_data, err := interactor.NewDepartamentoInteractor().Update(id, params)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Departamento actualizado con éxito", res_data)
}

func (ctrl departamentoController) Delete(c echo.Context) error {

	//get request data
	id := c.Param("id")
	//process
	res_data, err := interactor.NewDepartamentoInteractor().Delete(id)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Departamento borrado con éxito", res_data)
}

func (ctrl departamentoController) GetAllDepartamentoCiudad(c echo.Context) error {

	//get request data
	var params models.TbDepartamentoQuery
	if err := c.Bind(&params); err != nil {
		return res.RespError(c, err)
	}
	//validate request data
	if _, err_data := govalidator.ValidateStruct(params); err_data != nil {
		return res.RespError(c, err_data)
	}
	//process
	res_data, err := interactor.NewDepartamentoInteractor().GetAllDepartamentoCiudad(params)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Lista de departamentos y ciudades", res_data)
}
