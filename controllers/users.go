package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Greeting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GORU is working",
	})
}
