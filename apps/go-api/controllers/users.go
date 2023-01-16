package controllers

import (
	"net/http"

	"github.com/NegativeDevelopment/cerulean/go-api/lib"
	"github.com/NegativeDevelopment/cerulean/go-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// --- Routes ---
func UsersRoutes(parentGroup *gin.RouterGroup) {
	group := parentGroup.Group("/users")
	{
		group.GET("/:userid", getUser)
	}
}

// --- Handlers ---
func getUser(c *gin.Context) {
	userId := c.Param("userid")

	var user models.User
	if err := lib.DB.Select("id", "username").Where("id = ?", userId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
			return
		}
	}

	c.JSON(http.StatusOK, user)
}
