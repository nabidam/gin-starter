package panics

import (
	"github.com/nabidam/gin-starter/internal/utils"
)

func UserNotFound() {
	err := utils.UserNotFoundError()
	utils.ErrorPanic(err)
}
