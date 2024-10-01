package greetController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Greet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GORU is working",
	})
}
