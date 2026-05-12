package main

import (
	"github.com/MCH4X/wem-app/backend/internal/database"
	"github.com/MCH4X/wem-app/backend/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := database.InitDB()
	routes.SetupRouter(db)
}
