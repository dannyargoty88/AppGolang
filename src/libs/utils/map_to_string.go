package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func MapToString(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}

func InterfaceToString(m interface{}) string {
	json_data, _ := json.Marshal(m)
	return string(json_data)
}
