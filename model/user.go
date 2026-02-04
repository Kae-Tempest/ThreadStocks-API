package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string   `gorm:"unique" json:"username"`
	Password string   `json:"-"` // Password excluded from JSON responses
	Email    string   `gorm:"unique" json:"email"`
	Threads  []Thread `gorm:"foreignKey:UserID" json:"threads"`
}