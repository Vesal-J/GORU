package middleware

import (
	"github.com/gin-gonic/gin"
	"goru/response/errors"
	"goru/services/AuthService"
	"strings"
)

func UserAuthMiddleware(c *gin.Context) {

	token := c.GetHeader("Authorization")
	if token == "" {
		errors.UnauthorizedResponse(c)
		c.Abort()
		return
	}
	token = strings.Split(token, "Bearer ")[1]

	if _, err := AuthService.CheckToken(token); err != nil {
		errors.UnauthorizedResponse(c)
		c.Abort()
		return
	}
	c.Next()
}
