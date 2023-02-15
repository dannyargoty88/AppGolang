package models

type Email struct {
	To           string       `json:"to"`
	Addressee    string       `json:"addressee"`
	Subject      string       `json:"subject"`
	BodyData     interface{}  `json:"body"`
	BodyTemplate string       `json:"template"`
	Attach       []DataAttach `json:"attach"`
}

type DataAttach struct {
	File string `json:"file"`
}

type ConfigEmail struct {
	User       string `json:"user"`
	Password   string `json:"password"`
	Port       int    `json:"port"`
	Host       string `json:"host"`
	Smtpsecure string `json:"smtpsecure"`
}

type BodyEmailTest struct {
	Codigo  int
	Nombre  string
	Pokemon string
}
