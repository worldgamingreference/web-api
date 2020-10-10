package main

import (
	"log"

	models "github.com/Singebob/web-api/model"
	"github.com/Singebob/web-api/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := gin.Default()
	models.ConnectDataBase()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/servers", router.FindServers)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
