package routes

import (
	"go-posts/controller"
	"go-posts/database"
	"go-posts/repository"
	"go-posts/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	db := database.DB

	postRepo := repository.NewPostRepository(db)
	postService := services.NewPostService(postRepo)
	postController := controller.NewPostHandler(postService)

	r.GET("/articles/:limit/:offset", postController.GetPosts)
	r.GET("/article/:id", postController.GetPostByID)
	r.POST("article", postController.CreatePost)
	r.PUT("article/:id", postController.UpdatePost)
	r.DELETE("article/:id", postController.DeletePost)

}
