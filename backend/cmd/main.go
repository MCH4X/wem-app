package main

import (
	"github.com/MCH4X/wem-app/backend/internal/database"
	model "github.com/MCH4X/wem-app/backend/internal/models"
	"github.com/MCH4X/wem-app/backend/internal/routes"
)

func main() {
	db := database.InitDB()
	// auto migrate
	db.AutoMigrate(&model.User{})

	routes.SetupRouter(db)
}
