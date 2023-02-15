package repository

import (
	db "app/infrastructure/datastore"
	utils "app/src/libs/utils"
	models "app/src/models"
	"errors"

	"github.com/labstack/echo/v4"
)

//Builder
type errorRepository struct{}

func NewErrorRepository() ErrorRepository {
	return &errorRepository{}
}

type ErrorRepository interface {
	FindAll(params models.TbErrorQuery) ([]models.TbError, error)
	FindOne(id string) (*models.TbError, error)
	Create(report *echo.HTTPError, c echo.Context) (string, error)
	Delete(id string) (*models.TbError, error)
}

//Methods
func (repo errorRepository) FindAll(params models.TbErrorQuery) ([]models.TbError, error) {

	var data []models.TbError

	var query = db.Conn().Order("error_codigo desc")

	if params.Error_codigo > 0 {
		query = query.Where("error_codigo = ?", params.Error_codigo)
	}

	if params.Error_fecha != "" {
		query = query.Where("error_fecha = ?", params.Error_fecha)
	}

	if params.Error_url != "" {
		query = query.Where("error_url ilike ?", "%"+params.Error_url+"%")
	}

	if params.Error_message != "" {
		query = query.Where("error_message ilike ?", "%"+params.Error_message+"%")
	}

	err := query.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo errorRepository) FindOne(id string) (*models.TbError, error) {

	var data models.TbError

	err := db.Conn().Model(data).First(&data, []string{id}).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (repo errorRepository) Create(report *echo.HTTPError, c echo.Context) (string, error) {

	//crear log de errores en db
	var errModel models.TbError
	errModel.Error_fecha = "NOW()"
	errModel.Error_ip = c.Request().RemoteAddr
	errModel.Error_host = c.Request().Host
	errModel.Error_useragent = c.Request().UserAgent()
	errModel.Error_method = c.Request().Method
	errModel.Error_url = c.Request().URL.String()
	errModel.Error_status = utils.ConvertInt(report.Code)
	errModel.Error_message = report.Message.(string)

	if errCreate := db.Conn().Create(&errModel).Error; !errors.Is(errCreate, nil) {
		return "", errCreate
	}
	return utils.ConvertInt(errModel.Error_codigo), nil
}

func (repo errorRepository) Delete(id string) (*models.TbError, error) {

	var data models.TbError
	err := db.Conn().Where("error_codigo = ?", id).Delete(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
