package AuthService

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"goru/models"
	"time"
)

var env, _ = godotenv.Read(".env")

var secretKey []byte = []byte(env["JWT_SECRET_KEY"])

type GoruClaims struct {
	Id       int32
	Username string
	Exp      time.Time
	jwt.MapClaims
}

// takes a user and expiration in hours
func CreateToken(user *models.User, expiration int32) interface{} {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * time.Duration(expiration)),
	})

	plainTextToken, err := token.SignedString(secretKey)

	if err != nil {
		fmt.Println(err)
	}

	return plainTextToken
}

func CheckToken(stringToken string) (*GoruClaims, error) {

	token, err := jwt.ParseWithClaims(stringToken, &GoruClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*GoruClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("Token is not valid")
	}
}
