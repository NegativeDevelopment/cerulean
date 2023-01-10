package controllers

import (
	"github.com/NegativeDevelopment/cerulean/go-api/lib"
	"github.com/NegativeDevelopment/cerulean/go-api/models"
	"github.com/gin-gonic/gin"
)

func UserRoutes(parentGroup *gin.RouterGroup) {
	group := parentGroup.Group("/users")
	{
		group.GET("/:userid", getUser)
	}
}

func getUser(c *gin.Context) {
	userId := c.Param("userid")

	var user models.User
	if err := lib.DB.Select("id", "username").Where("id = ?", userId).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}
