package models

type Pdf struct {
	BodyPdf      interface{} `json:"body"`
	BodyTemplate string      `json:"template"`
	PrintPath    string      `json:"print_path"`
}

type BodyPdfTest struct {
	Codigo string `json:"codigo"`
	Nombre string `json:"nombre"`
	Valor  string `json:"valor"`
}
