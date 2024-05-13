package main

import (
	"fmt"
	userController "goru/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	init_db()
	router := gin.Default()

	router.GET("/", middleware, userController.Greeting)

	router.Run()
}

func middleware(c *gin.Context) {
	fmt.Println("Hello")
	c.Next()
}
