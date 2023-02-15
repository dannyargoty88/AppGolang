package presenter

import (
	models "app/src/models"
)

//Builder
type usuarioPresenter struct{}

func NewUsuarioPresenter() UsuarioPresenter {
	return &usuarioPresenter{}
}

type UsuarioPresenter interface {
	ResponseAll(data []models.TbUsuarioResponse) []models.TbUsuarioResponse
	ResponseOne(data *models.TbUsuarioResponse) *models.TbUsuarioResponse
	Response(data *models.TbUsuario) *models.TbUsuario
}

//Methods
func (prs usuarioPresenter) ResponseAll(data []models.TbUsuarioResponse) []models.TbUsuarioResponse {
	for i := range data {
		data[i].Usuario_password = "*******"
	}
	return data
}

func (prs usuarioPresenter) ResponseOne(data *models.TbUsuarioResponse) *models.TbUsuarioResponse {
	data.Usuario_password = "*******"
	return data
}

func (prs usuarioPresenter) Response(data *models.TbUsuario) *models.TbUsuario {
	data.Usuario_password = "*******"
	return data
}
