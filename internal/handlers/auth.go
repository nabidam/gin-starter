package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nabidam/gin-starter/internal/data/request"
	_ "github.com/nabidam/gin-starter/internal/data/request"
	_ "github.com/nabidam/gin-starter/internal/data/response"
	"github.com/nabidam/gin-starter/internal/repository"
)

type AuthHandler struct {
	authRepository repository.AuthRepository
}

func NewAuthHandler(authRepository repository.AuthRepository) *AuthHandler {
	return &AuthHandler{
		authRepository: authRepository,
	}
}

// @Summary Login using email and password
// @Schemes
// @Description Login using email and password
// @Tags auth
// @Produce json
// @Param user body request.UserLogin true "User data"
// @Success 200 {object} response.UserResponse
// @Failure 404 {object} response.UserResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var user request.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loggedInUser, err := h.authRepository.Login(user)
	if loggedInUser == nil {
		c.AbortWithStatus(404)
	}
	if err != nil {
		c.AbortWithError(401, err)
		return
	}

	userOut := loggedInUser.ToResponse()
	c.JSON(http.StatusOK, userOut)
}
