package repository

import (
	db "app/infrastructure/datastore"
	libsemail "app/src/libs/email"
	libsws "app/src/libs/request"
	utils "app/src/libs/utils"
	models "app/src/models"
	"encoding/json"
	"errors"
	"net/http"
)

//Builder
type terceroRepository struct{}

func NewTerceroRepository() TerceroRepository {
	return &terceroRepository{}
}

type TerceroRepository interface {
	FindAll(params models.TbTerceroQuery) ([]models.TbTerceroResponse, error)
	FindOne(id string) (*models.TbTerceroResponse, error)
	Create(data *models.TbTercero) (*models.TbTercero, error)
	Update(id string, data *models.TbTercero) (*models.TbTercero, error)
	Delete(id string) (*models.TbTercero, error)
	SendEmail(id string) (*models.TbTerceroEmailResponse, error)
}

//Methods
func (repo terceroRepository) FindAll(params models.TbTerceroQuery) ([]models.TbTerceroResponse, error) {

	var data []models.TbTerceroResponse

	var query = db.Conn().
		Select("t.*, c.ciudad_nombre, d.departamento_codigo, d.departamento_nombre").
		Table("tb_tercero t").
		Joins("JOIN tb_ciudad c ON c.ciudad_codigo = t.ciudad_codigo").
		Joins("JOIN tb_departamento d ON d.departamento_codigo = c.departamento_codigo").
		Order("t.tercero_codigo asc")

	if params.Tercero_codigo > 0 {
		query = query.Where("t.tercero_codigo = ?", params.Tercero_codigo)
	}

	if params.Tercero_documento != "" {
		query = query.Where("t.tercero_documento like ?", "%"+params.Tercero_documento+"%")
	}

	if params.Tercero_nombre != "" {
		query = query.Where("t.tercero_codigo ilike ?", "%"+params.Tercero_nombre+"%")
	}

	err := query.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo terceroRepository) FindOne(id string) (*models.TbTerceroResponse, error) {

	var data models.TbTerceroResponse

	var query = db.Conn().
		Select("t.*, c.ciudad_nombre, d.departamento_codigo, d.departamento_nombre").
		Table("tb_tercero t").
		Joins("JOIN tb_ciudad c ON c.ciudad_codigo = t.ciudad_codigo").
		Joins("JOIN tb_departamento d ON d.departamento_codigo = c.departamento_codigo").
		Where("t.tercero_codigo = ?", id)

	err := query.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (repo terceroRepository) Create(data *models.TbTercero) (*models.TbTercero, error) {

	if err := db.Conn().Create(data).Error; !errors.Is(err, nil) {
		return nil, err
	}
	return data, nil
}

func (repo terceroRepository) Update(id string, data *models.TbTercero) (*models.TbTercero, error) {

	err := db.Conn().Model(&data).Where("tercero_codigo = ?", id).Updates(&data).Error
	if err != nil {
		return nil, err
	}
	data.Tercero_codigo = utils.ConvertStr(id)
	return data, nil
}

func (repo terceroRepository) Delete(id string) (*models.TbTercero, error) {

	var data models.TbTercero
	err := db.Conn().Where("tercero_codigo = ?", id).Delete(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo terceroRepository) SendEmail(id string) (*models.TbTerceroEmailResponse, error) {

	var data models.TbTerceroEmailResponse

	//Obtener data del tercero
	tercero, err := repo.FindOne(id)
	if err != nil {
		return nil, err
	}

	//Consumir API REST
	data_ws, err := libsws.Request(models.RequestWs{
		Ws:       1,
		Method:   "GET",
		Endpoint: "https://pokeapi.co/api/v2/pokemon",
		Data:     nil,
		Headers: http.Header{
			"Content-Type": {"application/json; charset=utf-8"},
		},
	})
	if err != nil {
		return nil, err
	}
	response := models.Pokemon{}
	err = json.Unmarshal(data_ws, &response)
	if err != nil {
		return nil, err
	}

	//Documentos a enviar
	var attach []models.DataAttach
	attach = append(attach, models.DataAttach{File: "./infrastructure/storage/img/golang.png"})

	//Envio Email
	data_email := models.Email{}
	data_email.Subject = "Test send email"
	data_email.To = "correo_remitente@gmail.com"
	data_email.Addressee = "correo_destinatario@gmail.com"
	data_email.BodyData = models.BodyEmailTest{Codigo: tercero.Tercero_codigo, Nombre: tercero.Tercero_nombre, Pokemon: response.Results[5].Name}
	data_email.BodyTemplate = "./infrastructure/storage/template/email/email_test_template.html"
	data_email.Attach = attach

	err = libsemail.SendMail(data_email)
	if err != nil {
		return nil, err
	}

	//Respuesta Api
	data.Tercero_codigo = tercero.Tercero_codigo
	data.Tercero_nombre = tercero.Tercero_nombre
	data.Pokemon = response.Results[5].Name

	return &data, nil
}
