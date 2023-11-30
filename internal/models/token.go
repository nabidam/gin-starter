package models

import (
	"github.com/nabidam/gin-starter/internal/config"
	"github.com/nabidam/gin-starter/internal/data/response"
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	Token  string `gorm:"not null"`
	UserID uint64 `gorm:"index;not null"`
}

func init() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&Token{})
}

func (t *Token) ToResponse() response.TokenResponse {
	tokenResponse := response.TokenResponse{
		Token: t.Token,
	}
	return tokenResponse
}
