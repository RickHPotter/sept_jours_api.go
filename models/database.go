package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dataSourceName := os.Getenv("DB")

	if dataSourceName == "" {
		log.Fatal("No DB config found in .env file.")
	}

	DB, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		log.Panic("Failed to connect to DB.")
	}
}

func SyncDatabase() {
	DB.AutoMigrate(&(User{}))
	DB.AutoMigrate(&(DiaryEntry{}))
}
