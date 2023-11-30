package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/nabidam/gin-starter/internal/data/request"
	_ "github.com/nabidam/gin-starter/internal/data/response"
	"github.com/nabidam/gin-starter/internal/models"
	"github.com/nabidam/gin-starter/internal/repository"
	"github.com/nabidam/gin-starter/internal/utils/resource"
)

type UserHandler struct {
	userRepository repository.UserRepository
}

func NewUserHandler(userRepository repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepository: userRepository,
	}
}

// @Summary Get all users
// @Schemes
// @Description Get all users
// @Tags users
// @Produce json
// @Success 200 {array} response.UserResponse
// @Router /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	result := h.userRepository.GetAll()

	users := resource.ToResponse(result)

	c.JSON(http.StatusOK, users)
}

// @Summary Create new user
// @Schemes
// @Description Create new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body request.UserCreate true "User data"
// @Success 200 {object} response.UserResponse
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hash password
	user.HashPassword(user.Password)
	err := h.userRepository.Create(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userOut := user.ToResponse()
	c.JSON(http.StatusOK, userOut)
}

// @Summary Get logged in user
// @Schemes
// @Description Get logged in user
// @Tags users
// @Produce json
// @Security Bearer
// @Success 200 {object} response.UserResponse
// @Router /users/me [get]
func (h *UserHandler) GetMe(c *gin.Context) {
	authUser, _ := c.Get("user")
	user := authUser.(models.User)
	userOut := user.ToResponse()

	c.JSON(http.StatusOK, userOut)
}
