package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequestResponse(c *gin.Context, message string) {

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusBadRequest, gin.H{
		"message": message,
	})
}
