package controllers

import (
	"net/http"

	models "github.com/Singebob/web-api/model"
	"github.com/gin-gonic/gin"
)

// FindServers is function to get all servers
func FindServers(c *gin.Context) {
	var servers []models.Server
	models.DB.Find(&servers)
	c.JSON(http.StatusOK, servers)
}
