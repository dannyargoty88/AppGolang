package requestws

import (
	utils "app/src/libs/utils"
	models "app/src/models"
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

func Request(r models.RequestWs) ([]byte, error) {

	//Crear la peticón
	request, err := buildRequest(r)
	if err != nil {
		return nil, err
	}

	//Guardar log de la petición
	save_request, err := SaveRequest(r)
	if err != nil {
		return nil, err
	}

	//Ejecutar la petición
	response, err := customClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	//Leer respuesta
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	status_code := utils.ConvertInt(response.StatusCode)

	//Guardar log de la respuesta
	save_request.Itews_status = status_code
	save_request.Itews_response = string(body)
	err = SaveResponse(save_request)
	if err != nil {
		return nil, err
	}

	//Validar status de la respuesta
	if status_code[0:1] == "4" || status_code[0:1] == "5" {
		return nil, errors.New(string(body))
	}

	return body, nil
}

func buildRequest(r models.RequestWs) (*http.Request, error) {

	json_data, err := xml.Marshal(r.Data)
	request, err := http.NewRequest(r.Method, r.Endpoint, bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}
	request.Header = r.Headers

	return request, nil
}

func customClient() *http.Client {

	transport := http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{
		Transport: &transport,
	}

	return &client
}
