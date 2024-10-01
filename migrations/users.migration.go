package main

import (
	database "goru/db"
	models "goru/models"
)

func main() {
	if err := database.Db.Exec(`DO $$ BEGIN
		CREATE TYPE user_role AS ENUM ('admin', 'user');
	EXCEPTION
		WHEN duplicate_object THEN null;
	END $$;`).Error; err != nil {
		println(err)
	}
	database.Db.AutoMigrate(&models.User{})
}
