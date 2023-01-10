package controllers

import (
	"github.com/NegativeDevelopment/cerulean/go-api/lib"
	"github.com/NegativeDevelopment/cerulean/go-api/middlewares"
	"github.com/NegativeDevelopment/cerulean/go-api/models"
	"github.com/gin-gonic/gin"
)

func MeRoutes(parentGroup *gin.RouterGroup) {
	group := parentGroup.Group("/me", middlewares.JwtAuthMiddleware())
	{
		group.GET("", getMe)
	}
}

func getMe(c *gin.Context) {
	userId := c.MustGet("user").(string)

	var user models.User
	if err := lib.DB.Select("id", "username").Where("id = ?", userId).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}
