package greetController

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Greet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GORU is working",
	})
}
