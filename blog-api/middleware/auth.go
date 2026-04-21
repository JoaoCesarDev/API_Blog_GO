package middleware

import(
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("jcbest")

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenStr := strings.Split(c.GetHeader("Authorization"), "Bearer ")[1]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
		
	}
}