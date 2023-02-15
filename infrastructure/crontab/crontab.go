package crontab

import (
	clear "app/infrastructure/crontab/testcrontab"

	"github.com/jinzhu/gorm"
	"github.com/robfig/cron"
)

func Crontab(db *gorm.DB) {

	//Instancia de Cron
	c := cron.New()

	//Programar tareas ("second, minute, hour, dayofmonth, month, dayofweek")
	c.AddFunc("0 0 0 * * *", func() { clear.TestCrontab(db) })

	//Comenzar tareas
	c.Start()

	//Pausar o detener la ejecución
	//select {} //defer c.Stop()
}
