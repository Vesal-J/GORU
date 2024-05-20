package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
}

type APIUser struct {
	Username string
}
