package main

import "github.com/nabidam/gin-starter/internal/router"

func main() {
	// @securityDefinitions.apikey Bearer
	// @in header
	// @name Authorization
	// @description Type "Bearer" followed by a space and JWT token.
	r := router.SetupRouter()
	r.Run(":8000")
}
