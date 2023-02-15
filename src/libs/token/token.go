package token

import (
	config "app/infrastructure/config"
	models "app/src/models"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func GetSigningKey() []byte {
	config.ReadConfig()
	return []byte(config.C.JWT.SigningkeyApp)
}

func SigningMethod() *jwt.SigningMethodHMAC {

	config.ReadConfig()

	var signing_method *jwt.SigningMethodHMAC

	switch config.C.JWT.Signingmethod {
	case "HS256":
		signing_method = jwt.SigningMethodHS256
	default:
		signing_method = jwt.SigningMethodHS512
	}

	return signing_method
}

func CreateJwtToken(u models.TbUsuario) (string, error) {

	token := jwt.New(SigningMethod())

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["userCode"] = u.Usuario_codigo

	token_string, err := token.SignedString(GetSigningKey())
	if err != nil {
		return "", err
	}

	return token_string, nil
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {

	hmacSecretString := GetSigningKey()
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
