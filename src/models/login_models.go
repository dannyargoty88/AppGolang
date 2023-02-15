package models

//input struct
type Login struct {
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required"`
	Nombre   string `json:"nombre,omitempty"`
	Token    string `json:"token,omitempty"`
}

func (Login) TableName() string { return "public.usuario" }
