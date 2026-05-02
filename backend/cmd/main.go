package main

import (
	"github.com/MCH4X/wem-app/backend/internal/database"
	"github.com/MCH4X/wem-app/backend/internal/routes"
)

func main() {
	db := database.InitDB()
	routes.SetupRouter(db)
}
