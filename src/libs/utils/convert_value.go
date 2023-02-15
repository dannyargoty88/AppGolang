package utils

import (
	"strconv"
)

// Convertir parámetro string a int
func ConvertStr(param string) int {
	id, _ := strconv.Atoi(param)
	return id
}

// Convertir parámetro int a string
func ConvertInt(param int) string {
	num := strconv.Itoa(param)
	return num
}
