package models

//input struct
type TbWs struct {
	Ws_codigo            int    `json:"ws_codigo" gorm:"primary_key" valid:"required"`
	Tipws_codigo         int    `json:"tipws_codigo" gorm:"primary_key"`
	Ws_nombre            string `json:"ws_nombre" valid:"required"`
	Ws_url               string `json:"ws_url" valid:"required"`
	Usuario_codigo       int    `json:"usuario_codigo" valid:"required"`
	Ws_fechacreacion     string `json:"ws_fechacreacion"`
	Ws_fechamodificacion string `json:"ws_fechamodificacion"`
}

func (TbWs) TableName() string { return "public.tb_ws" }

//output struct
type TbWsResponse struct {
	Ws_codigo            int    `json:"ws_codigo"`
	Tipws_codigo         int    `json:"tipws_codigo"`
	Ws_nombre            string `json:"ws_nombre"`
	Ws_url               string `json:"ws_url"`
	Usuario_codigo       int    `json:"usuario_codigo"`
	Ws_fechacreacion     string `json:"ws_fechacreacion"`
	Ws_fechamodificacion string `json:"ws_fechamodificacion"`
}

func (TbWsResponse) TableName() string { return "public.tb_ws" }

//query struct
type TbWsQuery struct {
	Ws_codigo int    `json:"ws_codigo"`
	Ws_nombre string `json:"ws_nombre"`
	Ws_url    string `json:"ws_url"`
}
