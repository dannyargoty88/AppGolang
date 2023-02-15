package presenter

import (
	models "app/src/models"
)

//Builder
type ciudadPresenter struct{}

func NewCiudadPresenter() CiudadPresenter {
	return &ciudadPresenter{}
}

type CiudadPresenter interface {
	ResponseAll(data []models.TbCiudadResponse) []models.TbCiudadResponse
	ResponseOne(data *models.TbCiudadResponse) *models.TbCiudadResponse
	Response(data *models.TbCiudad) *models.TbCiudad
}

//Methods
func (prs ciudadPresenter) ResponseAll(data []models.TbCiudadResponse) []models.TbCiudadResponse {
	return data
}

func (prs ciudadPresenter) ResponseOne(data *models.TbCiudadResponse) *models.TbCiudadResponse {
	return data
}

func (prs ciudadPresenter) Response(data *models.TbCiudad) *models.TbCiudad {
	return data
}
