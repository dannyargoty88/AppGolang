package interactor

import (
	models "app/src/models"
	presenter "app/src/usecase/presenter"
	repository "app/src/usecase/repository"

	"github.com/labstack/echo/v4"
)

//Builder
type loginInteractor struct{}

func NewLoginInteractor() LoginInteractor {
	return &loginInteractor{}
}

type LoginInteractor interface {
	Login(params *models.Login, c echo.Context) (*models.Login, error)
}

//Methods
func (itr loginInteractor) Login(params *models.Login, c echo.Context) (*models.Login, error) {

	data, err := repository.NewLoginRepository().Login(params, c)
	if err != nil {
		return nil, err
	}
	return presenter.NewLoginInteractor().Response(data), nil
}
