package main

import (
	"go-posts/database"
	"go-posts/models"
	"log"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Connect()

	for i := 1; i <= 10; i++ {
		post := models.Post{
			Title:       "Post Title " + strconv.Itoa(i),
			Content:     "Lorem ipsum content " + strconv.Itoa(i),
			Category:    "general",
			Status:      "publish",
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		}
		database.DB.Create(&post)
	}

	log.Println("Seeded 10 posts")
}
