package repository

import (
	"github.com/nabidam/gin-starter/internal/models"
	"github.com/nabidam/gin-starter/internal/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetAll() []models.User
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetAll() []models.User {
	var users []models.User
	results := r.db.Find(&users)
	utils.ErrorPanic(results.Error)
	return users
}
