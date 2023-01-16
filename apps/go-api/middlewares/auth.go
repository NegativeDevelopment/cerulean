package middlewares

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Check if the authorization header is set
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "Missing auth token"})
			return
		}

		// Get the token from the header
		splitToken := strings.Split(tokenString, "Bearer ")
		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(401, gin.H{"message": "Invalid auth token"})
			return
		}

		// Parse the token
		token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "Invalid auth token"})
			return
		}

		// Check if the token is valid and set the user in the context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["userId"] == nil {
				c.AbortWithStatusJSON(401, gin.H{"message": "Invalid auth token"})
				return
			}

			userId, err := uuid.Parse(claims["userId"].(string))
			if err != nil {
				c.AbortWithStatusJSON(401, gin.H{"message": "Invalid auth token"})
				return
			}

			c.Set("user", userId)
		} else {
			c.AbortWithStatusJSON(401, gin.H{"message": "Invalid auth token"})
			return
		}

		c.Next()
	}
}
