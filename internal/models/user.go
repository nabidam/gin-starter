package models

import (
	"github.com/nabidam/gin-starter/internal/config"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
}

func init() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&User{})
}
