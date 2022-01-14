package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	db, err := gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})
	DB = db
}
