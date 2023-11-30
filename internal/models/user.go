package models

import (
	"github.com/nabidam/gin-starter/internal/config"
	"github.com/nabidam/gin-starter/internal/data/response"
	"github.com/nabidam/gin-starter/internal/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;index;not null"`
	Email    string `gorm:"unique;index;not null"`
	Password string `gorm:"not null"`
}

func init() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) ToResponse() response.UserResponse {
	userResponse := response.UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
	return userResponse
}

func (u *User) HashPassword(password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}

func (u *User) VerifyPassword(password string) bool {
	return utils.VerifyPassword(u.Password, password)
}
