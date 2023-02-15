package repository

import (
	db "app/infrastructure/datastore"
	models "app/src/models"
)

//Builder
type itemWsRepository struct{}

func NewItemWsRepository() ItemWsRepository {
	return &itemWsRepository{}
}

type ItemWsRepository interface {
	FindAll(params models.TbItemwsQuery) ([]models.TbItemwsResponse, error)
	FindOne(ws string, id string) (*models.TbItemwsResponse, error)
}

//Methods
func (repo itemWsRepository) FindAll(params models.TbItemwsQuery) ([]models.TbItemwsResponse, error) {

	var data []models.TbItemwsResponse

	var query = db.Conn().Order("itews_codigo asc")

	if params.Ws_codigo > 0 {
		query = query.Where("ws_codigo = ?", params.Ws_codigo)
	}

	if params.Itews_codigo > 0 {
		query = query.Where("itews_codigo = ?", params.Itews_codigo)
	}

	if params.Itews_endpoint != "" {
		query = query.Where("itews_endpoint = ?", params.Itews_endpoint)
	}

	err := query.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo itemWsRepository) FindOne(ws string, id string) (*models.TbItemwsResponse, error) {

	var data models.TbItemwsResponse

	err := db.Conn().Model(data).First(&data, []string{ws, id}).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}
