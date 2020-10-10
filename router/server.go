package router

import (
	"net/http"

	"github.com/Singebob/web-api/controllers"

	"github.com/gin-gonic/gin"
)

// FindServers is find on all server
func FindServers(c *gin.Context) {
	c.JSON(http.StatusOK, controllers.FindServers())
}
