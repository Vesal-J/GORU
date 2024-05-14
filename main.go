package main

import (
	userController "goru/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	connect_to_db()
	init_db()
	router := gin.Default()

	router.GET("/", userController.Greeting)

	router.Run()
}
