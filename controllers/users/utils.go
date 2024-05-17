package userController

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	models "goru/models"
	"time"
)

var env, _ = godotenv.Read(".env")

var secretKey []byte = []byte(env["JWT_SECRET_KEY"])

// takes a user and expiration in hours
func createToken(user *models.User, expiration int32) interface{} {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.UserName,
		"exp":      time.Now().Add(time.Hour * time.Duration(expiration)),
	})

	plainTextToken, err := token.SignedString(secretKey)

	if err != nil {
		fmt.Println(err)
	}

	return plainTextToken
}
