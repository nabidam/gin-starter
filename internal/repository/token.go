package repository

import (
	"github.com/nabidam/gin-starter/internal/models"
	"gorm.io/gorm"
)

type TokenRepository interface {
	CreateToken(token *models.Token) error
}

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{
		db: db,
	}
}

func (t *tokenRepository) CreateToken(token *models.Token) error {
	return t.db.Create(token).Error
}
