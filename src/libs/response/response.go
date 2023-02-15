package response

import (
	customErr "app/src/libs/error"
	models "app/src/models"
	repository "app/src/usecase/repository"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RespSuccess(c echo.Context, message string, data interface{}) error {

	//Crear respuesta
	response := models.SuccessResponse{
		Meta: models.MetaResponse{
			Success: true,
			Message: message,
		},
		Data: data,
	}

	//Response
	return c.JSON(http.StatusOK, response)
}

func RespError(c echo.Context, err error) error {

	//Obtener info del error
	var msgErrSave = ""
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	//Procesar el error
	res_err := customErr.ErrorProcess(report.Code, err.Error())

	//Guardar error en bd
	if res_err.ErrorSave {
		codErr, errPr := repository.NewErrorRepository().Create(report, c)
		if errPr != nil {
			log.Fatalln(errPr)
		}
		msgErrSave = ". Error #" + codErr
	}

	//crear respuesta
	response := models.ErrorResponse{
		Meta: models.MetaResponse{
			Success: false,
			Message: res_err.ErrorCodeText,
		},
		Error: res_err.ErrorText + msgErrSave,
	}

	//Response
	return c.JSON(res_err.ErrorCode, response)
}
