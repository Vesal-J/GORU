package models

import (
	"gorm.io/gorm"
)

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Role     Role `gorm:"type:user_role"`
	Mobile   string
}

type APIUser struct {
	Username string
}
