package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	_ = godotenv.Load()

	direction := "up"
	if len(os.Args) > 1 {
		direction = os.Args[1]
	}

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		getEnv("DB_USER", "postgres"),
		os.Getenv("DB_PASSWORD"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_NAME", "mydb"),
		getEnv("DB_SSLMODE", "disable"),
	)

	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		log.Fatal("failed to init migrations:", err)
	}
	defer func() {
		if srcErr, dbErr := m.Close(); srcErr != nil || dbErr != nil {
			log.Printf("migrate close errors: src=%v db=%v", srcErr, dbErr)
		}
	}()

	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("migration up failed:", err)
		}
		log.Println("migrations applied")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("migration down failed:", err)
		}
		log.Println("migrations rolled back")
	default:
		log.Fatalf("unknown direction %q — use 'up' or 'down'", direction)
	}
}
