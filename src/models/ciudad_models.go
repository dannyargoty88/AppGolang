package models

//input struct
type TbCiudad struct {
	Ciudad_codigo       int    `json:"ciudad_codigo" gorm:"primary_key"`
	Ciudad_nombre       string `json:"ciudad_nombre" valid:"required"`
	Departamento_codigo int    `json:"departamento_codigo" valid:"required"`
}

func (TbCiudad) TableName() string { return "public.tb_ciudad" }

//output struct
type TbCiudadResponse struct {
	Ciudad_codigo       int    `json:"ciudad_codigo" gorm:"primary_key"`
	Ciudad_nombre       string `json:"ciudad_nombre"`
	Departamento_codigo int    `json:"departamento_codigo,omitempty"`
	Departamento_nombre string `json:"departamento_nombre,omitempty"`
}

func (TbCiudadResponse) TableName() string { return "public.tb_ciudad" }

//query struct
type TbCiudadQuery struct {
	Ciudad_codigo       int    `json:"ciudad_codigo"`
	Ciudad_nombre       string `json:"ciudad_nombre"`
	Departamento_codigo int    `json:"departamento_codigo"`
}
