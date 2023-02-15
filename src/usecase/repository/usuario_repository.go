package repository

import (
	db "app/infrastructure/datastore"
	utils "app/src/libs/utils"
	models "app/src/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

//Builder
type usuarioRepository struct{}

func NewUsuarioRepository() UsuarioRepository {
	return &usuarioRepository{}
}

type UsuarioRepository interface {
	FindAll(params models.TbUsuarioQuery) ([]models.TbUsuarioResponse, error)
	FindOne(id string) (*models.TbUsuarioResponse, error)
	Create(data *models.TbUsuario) (*models.TbUsuario, error)
	Update(id string, data *models.TbUsuario) (*models.TbUsuario, error)
	Delete(id string) (*models.TbUsuario, error)
}

//Methods
func (repo usuarioRepository) FindAll(params models.TbUsuarioQuery) ([]models.TbUsuarioResponse, error) {

	var data []models.TbUsuarioResponse

	var query = db.Conn().Order("usuario_codigo asc")

	if params.Usuario_codigo > 0 {
		query = query.Where("usuario_codigo = ?", params.Usuario_codigo)
	}

	if params.Usuario_nombre != "" {
		query = query.Where("usuario_nombre ilike ?", "%"+params.Usuario_nombre+"%")
	}

	if params.Usuario_email != "" {
		query = query.Where("usuario_email ilike ?", "%"+params.Usuario_email+"%")
	}

	err := query.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo usuarioRepository) FindOne(id string) (*models.TbUsuarioResponse, error) {

	var data models.TbUsuarioResponse

	err := db.Conn().Model(data).First(&data, []string{id}).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (repo usuarioRepository) Create(data *models.TbUsuario) (*models.TbUsuario, error) {

	bcrypt_pass, err := bcrypt.GenerateFromPassword([]byte(data.Usuario_password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	data.Usuario_password = string(bcrypt_pass)

	if err := db.Conn().Create(data).Error; !errors.Is(err, nil) {
		return nil, err
	}
	return data, nil
}

func (repo usuarioRepository) Update(id string, data *models.TbUsuario) (*models.TbUsuario, error) {

	err := db.Conn().Model(&data).Where("usuario_codigo = ?", id).Updates(&data).Error
	if err != nil {
		return nil, err
	}
	data.Usuario_codigo = utils.ConvertStr(id)
	return data, nil
}

func (repo usuarioRepository) Delete(id string) (*models.TbUsuario, error) {

	var data models.TbUsuario
	err := db.Conn().Where("usuario_codigo = ?", id).Delete(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
