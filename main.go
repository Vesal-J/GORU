package main

import (
	greetController "goru/controllers/greet"
	userController "goru/controllers/users"
	db "goru/db"
	"goru/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect_to_db()
	// init_db()
	router := gin.Default()

	router.GET("/", middleware.UserAuthMiddleware, greetController.Greet)
	router.POST("/login", userController.Login)
	router.POST("/register", userController.Register)

	router.Run()
}
