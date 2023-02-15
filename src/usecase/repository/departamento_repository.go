package repository

import (
	db "app/infrastructure/datastore"
	utils "app/src/libs/utils"
	models "app/src/models"
	"errors"
)

//Builder
type departamentoRepository struct{}

func NewDepartamentoRepository() DepartamentoRepository {
	return &departamentoRepository{}
}

type DepartamentoRepository interface {
	FindAll(params models.TbDepartamentoQuery) ([]models.TbDepartamentoResponse, error)
	FindOne(id string) (*models.TbDepartamentoResponse, error)
	Create(data *models.TbDepartamento) (*models.TbDepartamento, error)
	Update(id string, data *models.TbDepartamento) (*models.TbDepartamento, error)
	Delete(id string) (*models.TbDepartamento, error)
	FindAllDepartamentoCiudad(params models.TbDepartamentoQuery) ([]models.TbDepartamentoCiudadResponse, error)
}

//Methods
func (repo departamentoRepository) FindAll(params models.TbDepartamentoQuery) ([]models.TbDepartamentoResponse, error) {

	var data []models.TbDepartamentoResponse

	var query = db.Conn().Order("departamento_codigo asc")

	if params.Departamento_codigo > 0 {
		query = query.Where("departamento_codigo = ?", params.Departamento_codigo)
	}

	if params.Departamento_nombre != "" {
		query = query.Where("departamento_nombre ilike ?", "%"+params.Departamento_nombre+"%")
	}

	err := query.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo departamentoRepository) FindOne(id string) (*models.TbDepartamentoResponse, error) {

	var data models.TbDepartamentoResponse

	err := db.Conn().Model(data).First(&data, []string{id}).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (repo departamentoRepository) Create(data *models.TbDepartamento) (*models.TbDepartamento, error) {

	if err := db.Conn().Create(data).Error; !errors.Is(err, nil) {
		return nil, err
	}
	return data, nil
}

func (repo departamentoRepository) Update(id string, data *models.TbDepartamento) (*models.TbDepartamento, error) {

	err := db.Conn().Model(&data).Where("departamento_codigo = ?", id).Updates(&data).Error
	if err != nil {
		return nil, err
	}
	data.Departamento_codigo = utils.ConvertStr(id)
	return data, nil
}

func (repo departamentoRepository) Delete(id string) (*models.TbDepartamento, error) {

	var data models.TbDepartamento
	err := db.Conn().Where("departamento_codigo = ?", id).Delete(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo departamentoRepository) FindAllDepartamentoCiudad(params models.TbDepartamentoQuery) ([]models.TbDepartamentoCiudadResponse, error) {

	var data []models.TbDepartamentoCiudadResponse

	var query = db.Conn().Order("departamento_codigo asc")

	if params.Departamento_codigo > 0 {
		query = query.Where("departamento_codigo = ?", params.Departamento_codigo)
	}

	if params.Departamento_nombre != "" {
		query = query.Where("departamento_nombre ilike ?", "%"+params.Departamento_nombre+"%")
	}

	err := query.Find(&data).Error
	if err != nil {
		return nil, err
	}

	for i := range data {

		var detalle_ciudad []models.TbCiudadResponse

		errDet := db.Conn().
			Select("ciudad_codigo, ciudad_nombre").
			Table("tb_ciudad").
			Where("departamento_codigo = ? ", data[i].Departamento_codigo).
			Find(&detalle_ciudad).Error
		if errDet != nil {
			return nil, errDet
		}

		data[i].Ciudades = detalle_ciudad
	}

	return data, nil
}
