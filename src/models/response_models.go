package models

type MetaResponse struct {
	Success bool   `json:"success" default:"true"`
	Message string `json:"message" default:"true"`
}

type SuccessResponse struct {
	Meta MetaResponse `json:"meta"`
	Data interface{}  `json:"data"`
}

type ErrorResponse struct {
	Meta  MetaResponse `json:"meta"`
	Error string       `json:"error"`
}
