package repository

import (
	db "app/infrastructure/datastore"
	utils "app/src/libs/utils"
	models "app/src/models"
	"errors"
)

//Builder
type ciudadRepository struct{}

func NewCiudadRepository() CiudadRepository {
	return &ciudadRepository{}
}

type CiudadRepository interface {
	FindAll(params models.TbCiudadQuery) ([]models.TbCiudadResponse, error)
	FindOne(id string) (*models.TbCiudadResponse, error)
	Create(data *models.TbCiudad) (*models.TbCiudad, error)
	Update(id string, data *models.TbCiudad) (*models.TbCiudad, error)
	Delete(id string) (*models.TbCiudad, error)
}

//Methods
func (repo ciudadRepository) FindAll(params models.TbCiudadQuery) ([]models.TbCiudadResponse, error) {

	var data []models.TbCiudadResponse

	var query = db.Conn().
		Select("c.*, d.departamento_nombre").
		Table("tb_ciudad c").
		Joins("JOIN tb_departamento d ON d.departamento_codigo = c.departamento_codigo").
		Order("c.ciudad_codigo asc")

	if params.Ciudad_codigo > 0 {
		query = query.Where("c.ciudad_codigo = ?", params.Ciudad_codigo)
	}

	if params.Ciudad_nombre != "" {
		query = query.Where("c.Ciudad_nombre ilike ?", "%"+params.Ciudad_nombre+"%")
	}

	if params.Departamento_codigo > 0 {
		query = query.Where("c.departamento_codigo = ?", params.Departamento_codigo)
	}

	err := query.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo ciudadRepository) FindOne(id string) (*models.TbCiudadResponse, error) {

	var data models.TbCiudadResponse

	var query = db.Conn().
		Select("c.*, d.departamento_nombre").
		Table("tb_ciudad c").
		Joins("JOIN tb_departamento d ON d.departamento_codigo = c.departamento_codigo").
		Where("c.ciudad_codigo = ?", id)

	err := query.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (repo ciudadRepository) Create(data *models.TbCiudad) (*models.TbCiudad, error) {

	if err := db.Conn().Create(data).Error; !errors.Is(err, nil) {
		return nil, err
	}
	return data, nil
}

func (repo ciudadRepository) Update(id string, data *models.TbCiudad) (*models.TbCiudad, error) {

	err := db.Conn().Model(&data).Where("ciudad_codigo = ?", id).Updates(&data).Error
	if err != nil {
		return nil, err
	}
	data.Ciudad_codigo = utils.ConvertStr(id)
	return data, nil
}

func (repo ciudadRepository) Delete(id string) (*models.TbCiudad, error) {

	var data models.TbCiudad
	err := db.Conn().Where("ciudad_codigo = ?", id).Delete(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
