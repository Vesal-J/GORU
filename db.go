package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn string = "root@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func connect_to_db() {
	if err != nil {
		fmt.Println("Couldn't connect to database!")
		fmt.Println()
	}
}
