package userController

import (
	"goru/models"
	"goru/response/errors"
	"goru/services/AuthService"
	"net/http"

	"github.com/gin-gonic/gin"
	database "goru/db"
)

func Login(c *gin.Context) {
	var requestData map[string]string

	var err = c.BindJSON(&requestData)

	if err != nil {
		errors.BadRequestResponse(c, "couldn't read your request payload")
	}

	var user models.User
	database.Db.Where("username = ?", requestData["username"]).First(&user)

	if user.Password == requestData["password"] {
		token := AuthService.CreateToken(&user, 24)

		c.JSON(http.StatusOK, gin.H{
			"user":  user,
			"token": token,
		})
		return
	}

	errors.UnauthorizedResponse(c)
}

func Register(c *gin.Context) {
	var requestData map[string]string
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		Username: requestData["username"],
		Password: requestData["password"],
	}

	database.Db.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"username": requestData["username"],
	})
}
