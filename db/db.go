package database

import (
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var envData, _ = godotenv.Read(".env")
var Db, err = gorm.Open(postgres.Open(envData["DATABASE_URL"]), &gorm.Config{})

func Connect_to_db() {
	if err != nil {
		fmt.Println("Couldn't connect to database!")
	}
}
