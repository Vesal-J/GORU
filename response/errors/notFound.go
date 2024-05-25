package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFoundResponse(c *gin.Context) {

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Not Found",
	})
}
