package controllers

import (
	res "app/src/libs/response"

	"github.com/labstack/echo/v4"
)

//Builder
type indexController struct{}

func NewIndexController() IndexController {
	return &indexController{}
}

type IndexController interface {
	Index(c echo.Context) error
}

//Methods
func (ctrl indexController) Index(c echo.Context) error {
	//response
	return res.RespSuccess(c, "Index", "App Golang")
}
