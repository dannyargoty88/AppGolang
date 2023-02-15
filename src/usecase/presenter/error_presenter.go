package presenter

import (
	models "app/src/models"
)

//Builder
type errorPresenter struct{}

func NewErrorPresenter() ErrorPresenter {
	return &errorPresenter{}
}

type ErrorPresenter interface {
	ResponseAll(data []models.TbError) []models.TbError
	ResponseOne(data *models.TbError) *models.TbError
	Response(data *models.TbError) *models.TbError
}

//Methods
func (prs errorPresenter) ResponseAll(data []models.TbError) []models.TbError {
	return data
}

func (prs errorPresenter) ResponseOne(data *models.TbError) *models.TbError {
	return data
}

func (prs errorPresenter) Response(data *models.TbError) *models.TbError {
	return data
}
