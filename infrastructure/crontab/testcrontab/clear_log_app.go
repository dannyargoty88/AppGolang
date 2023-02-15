package crontab

import (
	"log"

	"github.com/jinzhu/gorm"
)

func TestCrontab(db *gorm.DB) {
	log.Println("Execute Crontab: Test")
}
