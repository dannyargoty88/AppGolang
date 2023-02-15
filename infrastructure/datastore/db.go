package datastore

import (
	"app/infrastructure/config"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func NewDB() *gorm.DB {

	config.ReadConfig()

	var configData config.Database
	var dataBase string
	var err error
	var log_mode bool

	switch config.C.SERVERENV {
	case "test":
		configData = config.C.Database.Test
		log_mode = true
	case "pre-production":
		configData = config.C.Database.Preproduction
		log_mode = false
	case "production":
		configData = config.C.Database.Production
		log_mode = false
	default:
		configData = config.C.Database.Develoment
		log_mode = true
	}

	dataBase = fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		configData.Host,
		configData.Port,
		configData.User,
		configData.DBName,
		configData.Password,
		configData.Sslmode,
	)

	DB, err = gorm.Open(configData.DBType, dataBase)
	if err != nil {
		log.Panic("failed to connect database")
		log.Fatalln(err)
	}

	DB.LogMode(log_mode)

	return DB
}

func Conn() *gorm.DB {
	return DB
}
