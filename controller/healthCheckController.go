
package controller

import (
	"net/http"

	"github.com/gin-goinc/gin"
)

func HealthCheck(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}