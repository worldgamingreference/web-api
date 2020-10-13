package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB connection
var DB *gorm.DB

// ConnectDataBase is connect to database with var env
func ConnectDataBase() {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable host=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Server{})

	DB = db
}
