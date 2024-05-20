package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UnauthorizedResponse(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "unauthorized",
	})
}
