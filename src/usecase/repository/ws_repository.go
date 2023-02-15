package repository

import (
	db "app/infrastructure/datastore"
	utils "app/src/libs/utils"
	models "app/src/models"
	"errors"
)

//Builder
type wsRepository struct{}

func NewWsRepository() WsRepository {
	return &wsRepository{}
}

type WsRepository interface {
	FindAll(params models.TbWsQuery) ([]models.TbWsResponse, error)
	FindOne(id string) (*models.TbWsResponse, error)
	Create(data *models.TbWs) (*models.TbWs, error)
	Update(id string, data *models.TbWs) (*models.TbWs, error)
	Delete(id string) (*models.TbWs, error)
}

//Methods
func (repo wsRepository) FindAll(params models.TbWsQuery) ([]models.TbWsResponse, error) {

	var data []models.TbWsResponse

	var query = db.Conn().Order("ws_codigo asc")

	if params.Ws_codigo > 0 {
		query = query.Where("ws_codigo = ?", params.Ws_codigo)
	}

	err := query.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo wsRepository) FindOne(id string) (*models.TbWsResponse, error) {

	var data models.TbWsResponse

	err := db.Conn().Model(data).First(&data, []string{id}).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (repo wsRepository) Create(data *models.TbWs) (*models.TbWs, error) {

	data.Ws_fechacreacion = "CURRENT_DATE"
	data.Ws_fechamodificacion = "CURRENT_DATE"

	if err := db.Conn().Create(data).Error; !errors.Is(err, nil) {
		return nil, err
	}
	return data, nil
}

func (repo wsRepository) Update(id string, data *models.TbWs) (*models.TbWs, error) {

	data.Ws_fechamodificacion = "CURRENT_DATE"

	err := db.Conn().Model(&data).Where("ws_codigo = ?", id).Updates(&data).Error
	if err != nil {
		return nil, err
	}
	data.Ws_codigo = utils.ConvertStr(id)
	return data, nil
}

func (repo wsRepository) Delete(id string) (*models.TbWs, error) {

	var data models.TbWs
	err := db.Conn().Where("ws_codigo = ?", id).Delete(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
