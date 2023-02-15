package controllers

import (
	res "app/src/libs/response"
	models "app/src/models"
	interactor "app/src/usecase/interactor"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

//Builder
type loginController struct{}

func NewLoginController() LoginController {
	return &loginController{}
}

type LoginController interface {
	Login(c echo.Context) error
}

//Methods
func (ctrl loginController) Login(c echo.Context) error {

	//get request data
	var login_params *models.Login
	if err := c.Bind(&login_params); err != nil {
		return res.RespError(c, err)
	}
	//validate request data
	if _, err_data := govalidator.ValidateStruct(login_params); err_data != nil {
		return res.RespError(c, err_data)
	}
	//process
	login, err := interactor.NewLoginInteractor().Login(login_params, c)
	if err != nil {
		return res.RespError(c, err)
	}
	//response
	return res.RespSuccess(c, "Ingreso a la App", login)
}
