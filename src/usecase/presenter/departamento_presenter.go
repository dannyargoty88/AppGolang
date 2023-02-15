package presenter

import (
	models "app/src/models"
)

//Builder
type departamentoPresenter struct{}

func NewDepartamentoPresenter() DepartamentoPresenter {
	return &departamentoPresenter{}
}

type DepartamentoPresenter interface {
	ResponseAll(data []models.TbDepartamentoResponse) []models.TbDepartamentoResponse
	ResponseOne(data *models.TbDepartamentoResponse) *models.TbDepartamentoResponse
	Response(data *models.TbDepartamento) *models.TbDepartamento
	ResponseAllDepartamentoCiudad(data []models.TbDepartamentoCiudadResponse) []models.TbDepartamentoCiudadResponse
}

//Methods
func (prs departamentoPresenter) ResponseAll(data []models.TbDepartamentoResponse) []models.TbDepartamentoResponse {
	return data
}

func (prs departamentoPresenter) ResponseOne(data *models.TbDepartamentoResponse) *models.TbDepartamentoResponse {
	return data
}

func (prs departamentoPresenter) Response(data *models.TbDepartamento) *models.TbDepartamento {
	return data
}

func (prs departamentoPresenter) ResponseAllDepartamentoCiudad(data []models.TbDepartamentoCiudadResponse) []models.TbDepartamentoCiudadResponse {
	return data
}
