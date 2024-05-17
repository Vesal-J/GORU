package userController

import (
	"goru/models"
	"net/http"

	"github.com/gin-gonic/gin"
	database "goru/db"
)

func Login(c *gin.Context) {
	var user models.User
	database.Db.First(&user)
	token := createToken(&user, 24)
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}
