package models

import "time"

//input struct
type TbItemws struct {
	Ws_codigo               int       `json:"ws_codigo" gorm:"primary_key"`
	Itews_codigo            int       `json:"itews_codigo" gorm:"primary_key"`
	Itews_method            string    `json:"itews_method"`
	Itews_endpoint          string    `json:"itews_endpoint"`
	Itews_status            string    `json:"itews_status"`
	Itews_request           string    `json:"itews_request"`
	Itews_response          string    `json:"itews_response"`
	Itews_fecharequest      time.Time `json:"itews_fecharequest"`
	Itews_fecharesponse     time.Time `json:"itews_fecharesponse"`
	Itews_fechacreacion     time.Time `json:"itews_fechacreacion"`
	Itews_fechamodificacion time.Time `json:"itews_fechamodificacion"`
}

func (TbItemws) TableName() string { return "public.tb_itemws" }

//output struct
type TbItemwsResponse struct {
	Ws_codigo               int       `json:"ws_codigo"`
	Itews_codigo            int       `json:"itews_codigo"`
	Itews_method            string    `json:"itews_method"`
	Itews_endpoint          string    `json:"itews_endpoint"`
	Itews_status            string    `json:"itews_status"`
	Itews_request           string    `json:"itews_request"`
	Itews_response          string    `json:"itews_response"`
	Itews_fecharequest      time.Time `json:"itews_fecharequest"`
	Itews_fecharesponse     time.Time `json:"itews_fecharesponse"`
	Itews_fechacreacion     time.Time `json:"itews_fechacreacion"`
	Itews_fechamodificacion time.Time `json:"itews_fechamodificacion"`
}

func (TbItemwsResponse) TableName() string { return "public.tb_itemws" }

//query struct
type TbItemwsQuery struct {
	Ws_codigo      int    `json:"ws_codigo"`
	Itews_codigo   int    `json:"itews_codigo"`
	Itews_endpoint string `json:"itews_endpoint"`
}
