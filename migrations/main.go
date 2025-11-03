package main

import (
	"log"

	"go-posts/database"
	"go-posts/models"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Connect()

	err := database.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration success for posts table")
}
