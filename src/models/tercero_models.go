package models

//input struct
type TbTercero struct {
	Tercero_codigo    int    `json:"tercero_codigo" gorm:"primary_key"`
	Tercero_nombre    string `json:"tercero_nombre" valid:"required"`
	Tercero_documento string `json:"tercero_documento" valid:"required"`
	Ciudad_codigo     int    `json:"ciudad_codigo" valid:"required"`
}

func (TbTercero) TableName() string { return "public.tb_tercero" }

//output struct
type TbTerceroResponse struct {
	Tercero_codigo      int    `json:"tercero_codigo" gorm:"primary_key"`
	Tercero_nombre      string `json:"tercero_nombre"`
	Tercero_documento   string `json:"tercero_documento"`
	Ciudad_codigo       int    `json:"ciudad_codigo"`
	Ciudad_nombre       string `json:"ciudad_nombre"`
	Departamento_codigo int    `json:"departamento_codigo"`
	Departamento_nombre string `json:"departamento_nombre"`
}

type TbTerceroEmailResponse struct {
	Tercero_codigo int    `json:"tercero_codigo"`
	Tercero_nombre string `json:"tercero_nombre"`
	Pokemon        string `json:"pokemon"`
}

func (TbTerceroResponse) TableName() string { return "public.tb_tercero" }

//query struct
type TbTerceroQuery struct {
	Tercero_codigo      int    `json:"tercero_codigo"`
	Tercero_nombre      string `json:"tercero_nombre"`
	Tercero_documento   string `json:"tercero_documento"`
	Ciudad_codigo       int    `json:"ciudad_codigo"`
	Departamento_codigo int    `json:"departamento_codigo"`
}
