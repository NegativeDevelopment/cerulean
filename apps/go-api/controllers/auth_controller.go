package controllers

import (
	"github.com/NegativeDevelopment/cerulean/go-api/initializers"
	"github.com/NegativeDevelopment/cerulean/go-api/lib"
	"github.com/NegativeDevelopment/cerulean/go-api/models"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	authRoute := router.Group("/auth")
	{
		authRoute.POST("/login", login)
		authRoute.POST("/register", register)
	}
}

func register(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid JSON",
		})
		return
	}

	user.HashPassword()

	if result := initializers.DB.Create(&user); result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User created",
	})
}

func login(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid JSON",
		})
		return
	}

	var dbUser models.User
	if result := initializers.DB.Where("name = ?", user.Name).First(&dbUser); result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	if !dbUser.CheckPassword(user.Password) {
		c.JSON(400, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	tokenString, err := lib.GenerateJWTToken(&dbUser)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error while trying to generate a new authentication token",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User logged in",
		"token":   tokenString,
	})
}
