package presenter

import (
	models "app/src/models"
)

//Builder
type wsPresenter struct{}

func NewWsPresenter() WsPresenter {
	return &wsPresenter{}
}

type WsPresenter interface {
	ResponseAll(data []models.TbWsResponse) []models.TbWsResponse
	ResponseOne(data *models.TbWsResponse) *models.TbWsResponse
	Response(data *models.TbWs) *models.TbWs
}

//Methods
func (prs wsPresenter) ResponseAll(data []models.TbWsResponse) []models.TbWsResponse {
	return data
}

func (prs wsPresenter) ResponseOne(data *models.TbWsResponse) *models.TbWsResponse {
	return data
}

func (prs wsPresenter) Response(data *models.TbWs) *models.TbWs {
	return data
}
