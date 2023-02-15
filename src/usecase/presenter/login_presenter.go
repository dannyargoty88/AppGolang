package presenter

import (
	models "app/src/models"
)

//Builder
type loginPresenter struct{}

func NewLoginInteractor() LoginPresenter {
	return &loginPresenter{}
}

type LoginPresenter interface {
	Response(data *models.Login) *models.Login
}

//Methods
func (prs loginPresenter) Response(data *models.Login) *models.Login {
	data.Password = "*******"
	return data
}
