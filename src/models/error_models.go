package models

//input and output struct
type TbError struct {
	Error_codigo    int    `json:"error_codigo" gorm:"primary_key"`
	Error_fecha     string `json:"error_fecha"`
	Error_ip        string `json:"error_ip"`
	Error_host      string `json:"error_host"`
	Error_useragent string `json:"error_useragent"`
	Error_method    string `json:"error_method"`
	Error_url       string `json:"error_url"`
	Error_status    string `json:"error_status"`
	Error_message   string `json:"error_message"`
}

func (TbError) TableName() string { return "public.tb_error" }

//query struct
type TbErrorQuery struct {
	Error_codigo  int    `json:"error_codigo"`
	Error_fecha   string `json:"error_fecha"`
	Error_url     string `json:"error_url"`
	Error_message string `json:"error_message"`
}
