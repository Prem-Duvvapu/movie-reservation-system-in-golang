package main

import (
	"log"

	"movie-reservation-system/config"
	"movie-reservation-system/routes"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := config.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := routes.SetupRouter(db)
	log.Println("Server is running on http://localhost:8080") 
	router.Run(":8080")
}