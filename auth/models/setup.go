package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase(dbname string) {
	db, err := gorm.Open(sqlite.Open(dbname))
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})
	DB = db
}
