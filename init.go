package main

import (
	"goru/db"
	models "goru/models"
)

func init_db() {
	database.Db.AutoMigrate(&models.User{})
}
