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

	claims, err := AuthService.CheckToken(token)

	if err != nil {
		errors.UnauthorizedResponse(c)
		c.Abort()
		return
	}

	c.Set("userId", claims.Id)
	c.Next()
}
