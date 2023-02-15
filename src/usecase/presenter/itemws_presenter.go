package presenter

import (
	models "app/src/models"
)

//Builder
type itemWsPresenter struct{}

func NewItemWsPresenter() ItemWsPresenter {
	return &itemWsPresenter{}
}

type ItemWsPresenter interface {
	ResponseAll(data []models.TbItemwsResponse) []models.TbItemwsResponse
	ResponseOne(data *models.TbItemwsResponse) *models.TbItemwsResponse
	Response(data *models.TbItemws) *models.TbItemws
}

//Methods
func (prs itemWsPresenter) ResponseAll(data []models.TbItemwsResponse) []models.TbItemwsResponse {
	return data
}

func (prs itemWsPresenter) ResponseOne(data *models.TbItemwsResponse) *models.TbItemwsResponse {
	return data
}

func (prs itemWsPresenter) Response(data *models.TbItemws) *models.TbItemws {
	return data
}
