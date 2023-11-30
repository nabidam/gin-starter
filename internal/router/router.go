package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/nabidam/gin-starter/docs"
	"github.com/nabidam/gin-starter/internal/config"
	"github.com/nabidam/gin-starter/internal/handlers"
	"github.com/nabidam/gin-starter/internal/repository"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupRouter() *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	db = config.GetDB()

	// Initialize repositories and handlers
	userRepository := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepository)

	authRepository := repository.NewAuthRepository(db)
	authHandler := handlers.NewAuthHandler(authRepository)

	// Define routes
	api := r.Group("/api/v1")
	{
		usersRoute := api.Group("/users")
		{
			usersRoute.GET("/", userHandler.GetUsers)
			// A get function which returns a hello world string by json
			usersRoute.POST("/", userHandler.CreateUser)
			// Add more routes as needed
		}

		// auth
		authRoute := api.Group("/auth")
		{
			authRoute.POST("/login", authHandler.Login)
		}

	}

	// swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
