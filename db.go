package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var envData, _ = godotenv.Read(".env")
var db, err = gorm.Open(mysql.Open(envData["DATABASE_URL"]), &gorm.Config{})

func connect_to_db() {
	if err != nil {
		fmt.Println("Couldn't connect to database!")
	}
}
