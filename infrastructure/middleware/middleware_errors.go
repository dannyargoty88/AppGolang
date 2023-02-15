package middleware_app

import (
	res "app/src/libs/response"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	res.RespError(c, err)
}
