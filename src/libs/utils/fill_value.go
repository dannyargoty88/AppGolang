package utils

import (
	"fmt"
	"strings"
)

/**
Funcion que permite rellenar o completar una cadena con valores predeterminados hasta un limite establecido
	Ejemplo:  fillValue("r", "8", "X", "Test") Resultado: "TestXXXX"
	Ejemplo:  fillValue("l", "4", "0", "123") Resultado: "0000123"
*/

func FillValue(dir string, limit string, comodin string, text string) string {

	limitText := ConvertStr(limit)
	valueInt := ConvertStr(limit)
	length := len([]rune(text))
	repeat := 0

	if valueInt >= length {
		limitText = length
		repeat = valueInt - length
	}

	if dir == "l" {
		return fmt.Sprintf("%s"+text[0:limitText], strings.Repeat(comodin, repeat))
	} else {
		return fmt.Sprintf(text[0:limitText]+"%s", strings.Repeat(comodin, repeat))
	}
}
