package routes

import(
	"github.com/gin-gonic/gin"
	"blog-api/controllers"
	"blog-api/middleware"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/login", controllers.Login)

	protected := api.Group("/")

	protected.Use(middleware.AuthMiddleware())

	protected.POST("/posts", controllers.CreatePost)

	protected.GET("/posts", controllers.GetPosts)
}