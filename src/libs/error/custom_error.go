package customerror

import (
	"net/http"
	"strings"
)

type ErrorCustomProcess struct {
	ErrorCode     int
	ErrorCodeText string
	ErrorText     string
	ErrorSave     bool
}

func ErrorProcess(error_code int, error_text string) ErrorCustomProcess {

	//Se obtiene valores del error original
	res_error := ErrorCustomProcess{}
	err_cod := error_code
	err_msg := error_text
	err_save := false

	if error_text != "" {

		switch {

		//Errores de la peticion
		case strings.Index(err_msg, "Not Found") >= 0:
			err_msg = "Ruta no encontrada"
			err_cod = http.StatusNotFound

		case strings.Index(err_msg, "missing or malformed jwt") >= 0:
			err_msg = "No autorizado, se requiere iniciar sesión"
			err_cod = http.StatusUnauthorized

		case strings.Index(err_msg, "invalid or expired jwt") >= 0:
			err_msg = "No autorizado, token invalido o sesión caducada"
			err_cod = http.StatusUnauthorized

		case strings.Index(err_msg, "El usuario no existe") >= 0:
			err_msg = "No autorizado, el usuario no existe"
			err_cod = http.StatusUnauthorized

		case strings.Index(err_msg, "type error: expected") >= 0 ||
			strings.Index(err_msg, "looking for beginning of value") >= 0 ||
			strings.Index(err_msg, "invalid input syntax for type") >= 0 ||
			strings.Index(err_msg, "invalid character") >= 0:
			err_msg = "Se ha asignando un tipo de dato incorrecto en la estructura de la petición"
			err_cod = http.StatusBadRequest
			err_save = true

		case strings.Index(err_msg, "non zero value required") >= 0 ||
			strings.Index(err_msg, "does not validate as") >= 0:

			err_msg = strings.Replace(err_msg, "non zero value required", "campo requerido", -1)
			err_msg = strings.Replace(err_msg, "does not validate as", "no es válido como", -1)
			err_cod = http.StatusBadRequest

		case strings.Index(err_msg, "Request body can't be empty") >= 0:
			err_msg = "El cuerpo de la solicitud no puede estar vacío"
			err_cod = http.StatusBadRequest

		case strings.Index(err_msg, "function only accepts structs") >= 0:
			err_msg = "Se requiere definir una estructura de entrada"
			err_cod = http.StatusBadRequest

		case strings.Index(err_msg, "Method Not Allowed") >= 0:
			err_msg = "Se intenta usar un método no permitido"
			err_cod = http.StatusMethodNotAllowed
			err_save = true

		case strings.Index(err_msg, "The system cannot find the file specified") >= 0:
			err_msg = "El sistema no puede encontrar el archivo especificado"
			err_save = true

		case strings.Index(err_msg, "server is temporarily unable to service") >= 0 && strings.Index(err_msg, "Service Unavailable") >= 0:
			err_msg = "Servicio no disponible: el servidor no puede atender temporalmente"
			err_cod = http.StatusServiceUnavailable
			err_save = true

		//Errores en base de datos
		case strings.Index(err_msg, "record not found") >= 0:
			err_msg = "Registro no encontrado"
			err_cod = http.StatusOK
			err_save = true

		case strings.Index(err_msg, "violates foreign key") >= 0:
			err_msg = "Ha ocurrido un error de integridad referencial en base de datos"
			err_save = true

		case strings.Index(err_msg, "violates unique constraint") >= 0 ||
			strings.Index(err_msg, "violates check constraint") >= 0 ||
			strings.Index(err_msg, "duplicate key") >= 0:

			err_msg = "Ha ocurrido un error, registro duplicado o incumple una restricción única en base de datos"
			err_save = true

		case strings.Index(err_msg, "violates not-null constraint") >= 0:
			err_msg = "Ha ocurrido un error, se asigna un valor nulo a un campo obligatorio en base de datos"
			err_save = true

		case strings.Index(err_msg, "value too long for type character varying") >= 0:
			err_msg = "Ha ocurrido un error, se asigna un valor demasiado largo para un campo de texto en base de datos"
			err_save = true

		case strings.Index(err_msg, "but expression is of type") >= 0:
			err_msg = "Ha ocurrido un error, se asigna un tipo de valor incorrecto a un campo en base de datos"
			err_save = true

		case strings.Index(err_msg, "relation") >= 0 && strings.Index(err_msg, "already exist") >= 0:
			err_msg = "Ha ocurrido un error, se intenta implementar una relación ya existente en base de datos"
			err_save = true

		case strings.Index(err_msg, "column") >= 0 && strings.Index(err_msg, "does not exist") >= 0:
			err_msg = "Ha ocurrido un error, se intenta usar un campo no existente en base de datos"
			err_save = true

		case strings.Index(err_msg, "function") >= 0 && strings.Index(err_msg, "does not exist") >= 0:
			err_msg = "Ha ocurrido un error, se intenta usar una función no existente en base de datos"
			err_save = true

		case strings.Index(err_msg, "record") >= 0 && strings.Index(err_msg, "has no field") >= 0:
			err_msg = "Ha ocurrido un error, se intenta usar un registro no existente en base de datos"
			err_save = true

		case strings.Index(err_msg, "canceling statement due to user request") >= 0:
			err_msg = "Se ha cancelado la petición"
			err_save = true

		default:
			err_msg = strings.Replace(err_msg, "pq:", "", -1)
			err_save = true
		}
	}

	//Se registra datos del error procesado
	res_error.ErrorCode = err_cod
	res_error.ErrorCodeText = http.StatusText(err_cod)
	res_error.ErrorText = err_msg
	res_error.ErrorSave = err_save

	return res_error
}
