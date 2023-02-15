package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Database struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	Sslmode  string
	DBType   string
}

type Config struct {
	Database struct {
		Develoment    Database
		Test          Database
		Preproduction Database
		Production    Database
	}
	JWT struct {
		Signingmethod string
		SigningkeyApp string
	}
	Server struct {
		Address string
	}
	SERVERENV string
}

var C Config

func ReadConfig() {

	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join("./infrastructure/config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
