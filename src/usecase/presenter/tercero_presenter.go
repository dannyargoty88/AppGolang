package presenter

import (
	models "app/src/models"
)

//Builder
type terceroPresenter struct{}

func NewTerceroPresenter() TerceroPresenter {
	return &terceroPresenter{}
}

type TerceroPresenter interface {
	ResponseAll(data []models.TbTerceroResponse) []models.TbTerceroResponse
	ResponseOne(data *models.TbTerceroResponse) *models.TbTerceroResponse
	Response(data *models.TbTercero) *models.TbTercero
	ResponseEmail(data *models.TbTerceroEmailResponse) *models.TbTerceroEmailResponse
}

//Methods
func (prs terceroPresenter) ResponseAll(data []models.TbTerceroResponse) []models.TbTerceroResponse {
	return data
}

func (prs terceroPresenter) ResponseOne(data *models.TbTerceroResponse) *models.TbTerceroResponse {
	return data
}

func (prs terceroPresenter) Response(data *models.TbTercero) *models.TbTercero {
	return data
}

func (prs terceroPresenter) ResponseEmail(data *models.TbTerceroEmailResponse) *models.TbTerceroEmailResponse {
	return data
}
