package controllers

import (
	res "app/src/libs/response"
	models "app/src/models"
	interactor "app/src/usecase/interactor"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

//Builder
type ciudadController struct{}

func NewCiudadController() CiudadController {
	return &ciudadController{}
}

type CiudadController interface {
	GetAll(c echo.Context) error
	GetOne(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

//Methods
func (ctrl ciudadController) GetAll(c echo.Context) error {

	//get request data
	var params models.TbCiudadQuery
	if err := c.Bind(&params); err != nil {
		return res.RespError(c, err)
	}
	//validate request data
	if _, err_data := govalidator.ValidateStruct(params); err_data != nil {
		return res.RespError(c, err_data)
	}
	//process
	res_data, err := interactor.NewCiudadInteractor().GetAll(params)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Lista de ciudades", res_data)
}

func (ctrl ciudadController) GetOne(c echo.Context) error {

	//get request data
	id := c.Param("id")
	//process
	res_data, err := interactor.NewCiudadInteractor().GetOne(id)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Información de la ciudad", res_data)
}

func (ctrl ciudadController) Create(c echo.Context) error {

	//get request data
	var params *models.TbCiudad
	if err := c.Bind(&params); err != nil {
		return res.RespError(c, err)
	}
	//validate request data
	if _, err_data := govalidator.ValidateStruct(params); err_data != nil {
		return res.RespError(c, err_data)
	}
	//process
	res_data, err := interactor.NewCiudadInteractor().Create(params)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Ciudad creada con éxito", res_data)
}

func (ctrl ciudadController) Update(c echo.Context) error {

	//get request data
	var params *models.TbCiudad
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
	res_data, err := interactor.NewCiudadInteractor().Update(id, params)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Ciudad actualizada con éxito", res_data)
}

func (ctrl ciudadController) Delete(c echo.Context) error {

	//get request data
	id := c.Param("id")
	//process
	res_data, err := interactor.NewCiudadInteractor().Delete(id)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Ciudad borrada con éxito", res_data)
}
