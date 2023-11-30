package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nabidam/gin-starter/internal/data/request"
	_ "github.com/nabidam/gin-starter/internal/data/response"
	"github.com/nabidam/gin-starter/internal/repository"
	"github.com/nabidam/gin-starter/internal/utils"
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
		utils.AbortWithError(c, err, http.StatusBadRequest)
		return
	}

	token, err := h.authRepository.Login(user)
	if token == nil || err != nil {
		utils.AbortWithError(c, err, http.StatusForbidden)
		return
	}

	tokenOut := token.ToResponse()
	c.JSON(http.StatusOK, tokenOut)
}
