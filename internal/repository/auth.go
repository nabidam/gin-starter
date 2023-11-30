package repository

import (
	"log"

	"github.com/nabidam/gin-starter/internal/data/request"
	"github.com/nabidam/gin-starter/internal/models"
	"github.com/nabidam/gin-starter/internal/utils"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(user request.UserLogin) (*models.Token, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Login(user request.UserLogin) (*models.Token, error) {
	userRepository := NewUserRepository(r.db)
	requestedUser, err := userRepository.FindByEmail(user.Email)

	if err != nil {
		return nil, err
	}

	log.Printf("User found: %s", requestedUser.Username)

	if requestedUser.VerifyPassword(user.Password) {
		// generate token
		tokenString, _ := utils.GenerateToken(requestedUser.ID)
		token := models.Token{
			Token:  tokenString,
			UserID: uint64(requestedUser.ID),
		}

		tokenRepository := NewTokenRepository(r.db)
		err := tokenRepository.CreateToken(&token)
		if err != nil {
			return &token, err
		}
		return &token, nil
	}
	return nil, utils.InvalidCredentialsError()
}
