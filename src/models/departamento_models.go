package models

//input struct
type TbDepartamento struct {
	Departamento_codigo int    `json:"departamento_codigo" gorm:"primary_key"`
	Departamento_nombre string `json:"departamento_nombre" valid:"required"`
}

func (TbDepartamento) TableName() string { return "public.tb_departamento" }

//output struct
type TbDepartamentoResponse struct {
	Departamento_codigo int    `json:"departamento_codigo" gorm:"primary_key"`
	Departamento_nombre string `json:"departamento_nombre"`
}

func (TbDepartamentoResponse) TableName() string { return "public.tb_departamento" }

type TbDepartamentoCiudadResponse struct {
	Departamento_codigo int                `json:"departamento_codigo" gorm:"primary_key"`
	Departamento_nombre string             `json:"departamento_nombre"`
	Ciudades            []TbCiudadResponse `json:"ciudades"`
}

func (TbDepartamentoCiudadResponse) TableName() string { return "public.tb_departamento" }

//query struct
type TbDepartamentoQuery struct {
	Departamento_codigo int    `json:"departamento_codigo"`
	Departamento_nombre string `json:"departamento_nombre"`
}
