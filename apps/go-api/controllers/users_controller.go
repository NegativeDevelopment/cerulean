package controllers

import (
	"github.com/NegativeDevelopment/cerulean/go-api/initializers"
	"github.com/NegativeDevelopment/cerulean/go-api/models"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	userRoute := router.Group("/users")
	{
		userRoute.GET("/:userid", getUser)
		userRoute.POST("/", createUser)
	}
}

func getUser(c *gin.Context) {

	userId := c.Param("userid")

	var user models.User
	result := initializers.DB.Where("id = ?", userId).First(&user)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": "User not found",
		})
		return
	}

	c.JSON(200, user)
}

func createUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid JSON",
		})
		return
	}

	initializers.DB.Create(&user)

	c.JSON(200, user)
}
