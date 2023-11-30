package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func AbortWithError(c *gin.Context, err error, status int) {
	c.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
}

func UserNotFoundError() error {
	err := errors.New("user not found")
	return err
}

func InvalidCredentialsError() error {
	err := errors.New("credentials are invalid")
	return err
}
