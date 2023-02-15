package repository

import (
	db "app/infrastructure/datastore"
	token "app/src/libs/token"
	models "app/src/models"
	"errors"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

//Builder
type loginRepository struct{}

func NewLoginRepository() LoginRepository {
	return &loginRepository{}
}

type LoginRepository interface {
	Login(params *models.Login, c echo.Context) (*models.Login, error)
}

//Methods
func (repo loginRepository) Login(login *models.Login, c echo.Context) (*models.Login, error) {

	var usuario models.TbUsuario

	//Buscar el correo del usuario en bd
	err := db.Conn().Find(&usuario, "usuario_email = ?", login.Email).Error
	if err != nil {
		return nil, errors.New("El usuario no existe")
	}

	//Comparar contraseña
	err_pwd := bcrypt.CompareHashAndPassword([]byte(usuario.Usuario_password), []byte(login.Password))
	if err_pwd != nil {
		return nil, errors.New("Credenciales incorrectas")
	}

	//Generar token
	v_token, err := token.CreateJwtToken(usuario)
	if err != nil {
		return nil, err
	}
	login.Token = v_token

	return login, nil
}
