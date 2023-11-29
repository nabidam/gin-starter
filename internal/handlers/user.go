package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nabidam/gin-starter/internal/models"
	"github.com/nabidam/gin-starter/internal/repository"
)

type UserHandler struct {
	userRepository repository.UserRepository
}

func NewUserHandler(userRepository repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepository: userRepository,
	}
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userRepository.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	result := h.userRepository.GetAll()

	// var users []response.UserResponse
	// for _, value := range result {
	// 	user := response.UserResponse{
	// 		ID:       value.ID,
	// 		Username: value.Username,
	// 		Email:    value.Email,
	// 		Password: value.Password,
	// 	}
	// 	users = append(users, user)
	// }

	c.JSON(http.StatusOK, result)
}
