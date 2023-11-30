package resource

import (
	"github.com/nabidam/gin-starter/internal/data/response"
	"github.com/nabidam/gin-starter/internal/models"
)

func ToResponse(users []models.User) []response.UserResponse {
	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, user.ToResponse())
	}

	return userResponses
}
