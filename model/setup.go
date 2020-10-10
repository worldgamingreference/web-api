package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB connection
var DB *gorm.DB

// ConnectDataBase is connect to database with var env
func ConnectDataBase() {
	dsn := "user=postgres password=pwd dbname=postgres port=5432 sslmode=disable host=localhost"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Server{})

	DB = db
}
