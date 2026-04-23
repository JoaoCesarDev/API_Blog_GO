package controllers

import (
	"blog-api/models"
	"blog-api/utils"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func Login(userCollection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req models.User

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "dados inválidos"})
			return
		}

		var user models.User

		err := userCollection.FindOne(
			context.TODO(),
			bson.M{"email": req.Email},
		).Decode(&user)
/*
		if err != nil {
			c.JSON(401, gin.H{"error": "credenciais inválidas"})
			return
		}
*/
		if !utils.CheckPasswordHash(req.Password, user.Password) {
			c.JSON(401, gin.H{"error": "credenciais inválidas"})
			return
		}

		token, err := utils.GenerateToken(user.ID.Hex())
		if err != nil {
			c.JSON(500, gin.H{"error": "erro ao gerar token"})
			return
		}

		c.JSON(200, gin.H{
			"token": token,
		})
	}
}

func Register(userCollection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req models.User

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "dados inválidos"})
			return
		}

		hash, _ := utils.HashPassword(req.Password)
		req.Password = hash

		result, err := userCollection.InsertOne(context.TODO(), req)
		if err != nil {
			c.JSON(500, gin.H{"error": "erro ao salvar"})
			return
		}

		c.JSON(201, gin.H{
			"id": result.InsertedID,
		})
	}
}