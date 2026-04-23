package routes

import (
	"blog-api/controllers"
	"blog-api/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	protected := api.Group("/")

	protected.Use(middleware.AuthMiddleware())

	protected.POST("/posts", controllers.CreatePost)

	protected.GET("/posts", controllers.GetPosts)
}

func AuthRoutes(r *gin.Engine,userCollection *mongo.Collection) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register(userCollection))
		auth.POST("/login", controllers.Login(userCollection))
	}
}