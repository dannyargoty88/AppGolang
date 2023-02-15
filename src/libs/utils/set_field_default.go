package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

/**
Funcion que permite establecer valores predeterminados a un struct
	Ejemplo:
		//definicion en estructura:
			type Prueba struct {
				Dato  string `json:"dato" limit:"1" default:"ABCD"`
			}
		//uso:
			myModel := &model.Prueba{}
			utils.Set(myModel, "default")
			fmt.Println(myModel.Dato)
		//resultado:
			"ABCD"
*/

func Set(ptr interface{}, tag string) error {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		return fmt.Errorf("Not a pointer")
	}

	v := reflect.ValueOf(ptr).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		if defaultVal := t.Field(i).Tag.Get(tag); defaultVal != "-" {
			if err := setField(v.Field(i), defaultVal); err != nil {
				return err
			}

		}
	}
	return nil
}

func setField(field reflect.Value, defaultVal string) error {

	if !field.CanSet() {
		return fmt.Errorf("Can't set value\n")
	}

	switch field.Kind() {

	case reflect.Int:
		if val, err := strconv.ParseInt(defaultVal, 10, 64); err == nil {
			field.Set(reflect.ValueOf(int(val)).Convert(field.Type()))
		}
	case reflect.String:
		field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))
	}

	return nil
}
