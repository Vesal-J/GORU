package main

import (
	models "goru/models"
)

func init_db() {
	db.AutoMigrate(&models.User{})
}
