package repository

import (
	"github.com/nabidam/gin-starter/internal/models"
	"github.com/nabidam/gin-starter/internal/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetAll() []models.User
	FindByUsername(username string) (models.User, error)
	FindByEmail(email string) (models.User, error)
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

func (r *userRepository) FindByUsername(username string) (models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
