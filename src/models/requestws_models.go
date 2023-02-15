package models

import "net/http"

type RequestWs struct {
	Ws       int         `json:"ws"`
	Method   string      `json:"method"`
	Endpoint string      `json:"endpoint"`
	Data     interface{} `json:"data"`
	Headers  http.Header `json:"headers"`
	Token    string      `json:"token"`
}
