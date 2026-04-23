package main

import (
	"github.com/gin-gonic/gin"
	"blog-api/config"
	"blog-api/routes"

)

func main() {
	config.ConnectDB()

	userCollection := config.DB.Collection("users")

	r := gin.Default()

	routes.SetupRoutes(r)

	routes.AuthRoutes(r, userCollection)

	r.Run(":8080")
}