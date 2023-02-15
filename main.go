package main

import (
	crontab "app/infrastructure/crontab"
	datastore "app/infrastructure/datastore"
	middleware "app/infrastructure/middleware"
	routes "app/infrastructure/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	//Iniciar conexión a base de datos
	db := datastore.NewDB()
	defer db.Close()

	//Iniciar instancia de la app
	e := echo.New()

	//Implementación de los middleware
	middleware.MiddlewareApp(e)

	//Llamado a las rutas de la app
	routes.Routes(e)

	//Ejecución de las tareas automáticas
	crontab.Crontab(db)

	//Ejecutar la app
	e.Logger.Fatal(e.Start(":2023"))
}
