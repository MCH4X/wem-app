package routes

import (
	"github.com/MCH4X/wem-app/backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.GET("/", handlers.HelloHandler)
}
