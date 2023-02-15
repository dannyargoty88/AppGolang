package email

import (
	utils "app/src/libs/utils"
	models "app/src/models"
	"bytes"
	"crypto/tls"
	"text/template"

	"gopkg.in/gomail.v2"
)

func SendMail(data_email models.Email) error {

	//Procesar el cuerpo y destinatarios del correo
	message, err := processBody(data_email)
	if err != nil {
		return err
	}

	//procesar lista de destinatarios
	listAdress := utils.Explode(",", data_email.Addressee)

	//Crear cliente email
	config_email := ConfigEmail()
	mail := gomail.NewDialer(config_email.Host, config_email.Port, config_email.User, config_email.Password)
	mail.TLSConfig = &tls.Config{InsecureSkipVerify: true, ServerName: config_email.Host}

	//Configuracion del envio del correo
	msg := gomail.NewMessage()
	msg.SetHeader("From", data_email.To)
	msg.SetHeader("To", listAdress...)
	msg.SetHeader("Subject", data_email.Subject)
	msg.SetBody("text/html", message)

	for i := range data_email.Attach {
		msg.Attach(data_email.Attach[i].File)
	}

	//Realizar envio de correo
	if err := mail.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}

//Funcion que asigna los valores de un Struct a un template HTML retornando un string
func processBody(data_email models.Email) (string, error) {

	t, err := template.ParseFiles(data_email.BodyTemplate)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, data_email.BodyData)
	if err != nil {
		return "", err
	}

	message := buf.String()

	return message, nil
}
