package email

import (
	"app/infrastructure/config"
	models "app/src/models"
)

func ConfigEmail() models.ConfigEmail {

	var config_email models.ConfigEmail

	config.ReadConfig()

	if config.C.SERVERENV == "production" {
		config_email.Host = "host"
		config_email.Port = 0000
		config_email.User = "user"
		config_email.Password = "password"
	} else {
		config_email.Host = "smtp.mailtrap.io"
		config_email.Port = 2525
		config_email.User = "a90e628d74a865"
		config_email.Password = "621bd90e388327"
	}

	return config_email
}
