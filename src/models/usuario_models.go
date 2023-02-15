package models

//input struct
type TbUsuario struct {
	Usuario_codigo   int    `json:"usuario_codigo" gorm:"primary_key"`
	Usuario_nombre   string `json:"usuario_nombre" valid:"required"`
	Usuario_email    string `json:"usuario_email" valid:"required, email"`
	Usuario_password string `json:"usuario_password" valid:"required"`
}

func (TbUsuario) TableName() string { return "public.tb_usuario" }

//output struct
type TbUsuarioResponse struct {
	Usuario_codigo   int    `json:"usuario_codigo" gorm:"primary_key"`
	Usuario_nombre   string `json:"usuario_nombre"`
	Usuario_email    string `json:"usuario_email"`
	Usuario_password string `json:"usuario_password"`
}

func (TbUsuarioResponse) TableName() string { return "public.tb_usuario" }

//query struct
type TbUsuarioQuery struct {
	Usuario_codigo int    `json:"usuario_codigo"`
	Usuario_nombre string `json:"usuario_nombre"`
	Usuario_email  string `json:"usuario_email"`
}
