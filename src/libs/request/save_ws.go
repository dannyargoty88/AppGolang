package requestws

import (
	db "app/infrastructure/datastore"
	"app/src/libs/utils"
	models "app/src/models"
	"errors"
	"time"
)

//Guardar Petición
func SaveRequest(r models.RequestWs) (*models.TbItemws, error) {

	var data models.TbItemws
	data.Ws_codigo = r.Ws
	data.Itews_method = r.Method
	data.Itews_endpoint = r.Endpoint
	data.Itews_request = utils.InterfaceToString(r.Data)
	data.Itews_fecharequest = time.Now()
	data.Itews_fechacreacion = time.Now()
	data.Itews_fechamodificacion = time.Now()

	if err := db.Conn().Create(&data).Error; !errors.Is(err, nil) {
		return nil, err
	}

	var fillData *models.TbItemws = &models.TbItemws{}
	err := db.Conn().Model(fillData).First(&fillData, data).Error
	if err != nil {
		return nil, err
	}

	return fillData, nil
}

//Guardar Respuesta
func SaveResponse(data *models.TbItemws) error {

	data.Itews_fecharesponse = time.Now()
	data.Itews_fechamodificacion = time.Now()

	err := db.Conn().Model(&data).
		Where("ws_codigo = ?", data.Ws_codigo).
		Where("itews_codigo = ?", data.Itews_codigo).
		Updates(&data).Error
	if err != nil {
		return err
	}

	return nil
}
