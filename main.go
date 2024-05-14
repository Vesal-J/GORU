package main

import (
	userController "goru/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	init_db()
	connect_to_db()
	router := gin.Default()

	router.GET("/", userController.Greeting)

	router.Run()
}
