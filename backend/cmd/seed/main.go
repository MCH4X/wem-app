package main

import (
	"log"

	"github.com/MCH4X/wem-app/backend/internal/database"
	model "github.com/MCH4X/wem-app/backend/internal/models"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := database.InitDB()

	var count int64
	db.Model(&model.User{}).Count(&count)
	if count > 0 {
		log.Println("database already seeded, skipping")
		return
	}

	users := []model.User{
		{Name: "Alice Johnson", Age: 28},
		{Name: "Bob Smith", Age: 35},
		{Name: "Charlie Brown", Age: 22},
		{Name: "Diana Prince", Age: 31},
		{Name: "Ethan Hunt", Age: 40},
	}

	if err := db.Create(&users).Error; err != nil {
		log.Fatal("seed failed:", err)
	}

	log.Printf("seeded %d users", len(users))
}
