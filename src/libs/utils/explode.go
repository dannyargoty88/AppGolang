package utils

import "strings"

/**
Funcion que recibe un delimitador y una cadena y devuelve un arreglo de string
	Ejemplo:
		Saludos := Explode(",", "Hola,Hello,Hi")
		for _, saludo := range Saludos {
			fmt.Println(saludo)
		}
		//resultado:
			Hola
			Hello
			Hi
*/

func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}
