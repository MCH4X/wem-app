package main

import (
	"github.com/MCH4X/wem-app/backend/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetupRouter(router)

	router.Run(":8080")
}
