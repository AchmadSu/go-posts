package main

import (
	"fmt"
	"go-posts/config"
	"go-posts/database"
	"go-posts/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.ConnectDB()

	r := gin.Default()
	r.Use(cors.Default())
	routes.SetupRoutes(r)

	fmt.Println("Server start on :8080")
	r.Run(":8080")
}
