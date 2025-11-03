package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") +
		"@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME") +
		"?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	DB = db
	log.Println("Database connected")
}
