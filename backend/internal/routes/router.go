package routes

import (
	"log"
	"os"

	handler "github.com/MCH4X/wem-app/backend/internal/handlers"
	repository "github.com/MCH4X/wem-app/backend/internal/repositories"
	service "github.com/MCH4X/wem-app/backend/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) {
	router := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.GET("/users", userHandler.GetUsers)
	router.POST("/users", userHandler.CreateUser)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
