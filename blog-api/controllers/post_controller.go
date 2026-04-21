package controllers

import (
	"blog-api/models"
	"blog-api/utils"
	"context"
	"net/http"
	"blog-api/config"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	token, _ := utils.GenerateToken("123")

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func CreatePost(c *gin.Context) {
	var post models.Post
	c.BindJSON(&post)

	collection := config.DB.Collection("posts")
	result, _ := collection.InsertOne(context.TODO(), post)

	c.JSON(http.StatusCreated,result)
}

func GetPosts(c *gin.Context) {
	collection := config.DB.Collection("posts")
	cursor, _ := collection.Find(context.TODO(), map[string]interface{}{})
	var posts []models.Post

	cursor.All(context.TODO(), &posts)

	c.JSON(http.StatusOK, posts)
}

