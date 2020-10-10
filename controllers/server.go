package controllers

import (
	models "github.com/Singebob/web-api/model"
)

// FindServers is function to get all servers
func FindServers() []models.Server {
	var servers []models.Server
	models.DB.Find(&servers)
	return servers
}
