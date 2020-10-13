package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB connection
var DB *gorm.DB

// ConnectDataBase is connect to database with var env
func ConnectDataBase() {
	dsn := "user=postgres password=pwd dbname=postgres port=5432 sslmode=disable host=localhost"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Server{})

	DB = db
}
