package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UnauthorizedResponse(c *gin.Context) {

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "unauthorized",
	})
}
