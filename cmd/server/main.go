package main

import "github.com/nabidam/gin-starter/internal/router"

func main() {
	r := router.SetupRouter()
	r.Run(":8000")
}
