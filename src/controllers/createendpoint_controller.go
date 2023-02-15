package controllers

import (
	res "app/src/libs/response"
	models "app/src/models"
	interactor "app/src/usecase/interactor"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

//Builder
type createEndpointController struct{}

func NewCreateEndpointController() CreateEndpointController {
	return &createEndpointController{}
}

type CreateEndpointController interface {
	Create(c echo.Context) error
}

//Methods
func (ctrl createEndpointController) Create(c echo.Context) error {

	//get request data
	var params *models.CreateEndpoint
	if err := c.Bind(&params); err != nil {
		return res.RespError(c, err)
	}
	//validate request data
	if _, err_data := govalidator.ValidateStruct(params); err_data != nil {
		return res.RespError(c, err_data)
	}
	//process
	res_data, err := interactor.NewCreateEndpointInteractor().Create(params)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Endpoint creado con éxito", res_data)
}
