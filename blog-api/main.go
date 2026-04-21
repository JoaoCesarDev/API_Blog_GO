package main

import (
	"github.com/gin-gonic/gin"
	"blog-api/config"
	"blog-api/routes"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}