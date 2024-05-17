package main

import (
	"github.com/gin-gonic/gin"
	greetController "goru/controllers/greet"
	userController "goru/controllers/users"
	db "goru/db"
)

func main() {
	db.Connect_to_db()
	init_db()
	router := gin.Default()

	router.GET("/", greetController.Greet)
	router.POST("/login", userController.Login)

	router.Run()
}
