package controllers

import (
	"net/http"

	"github.com/NegativeDevelopment/cerulean/go-api/lib"
	"github.com/NegativeDevelopment/cerulean/go-api/middlewares"
	"github.com/NegativeDevelopment/cerulean/go-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// --- Routes ---
func MyRoutes(parentGroup *gin.RouterGroup) {
	group := parentGroup.Group("/me", middlewares.JwtAuthMiddleware())
	{
		group.GET("", getMe)
	}
}

// --- Handlers ---
func getMe(c *gin.Context) {
	userId := c.MustGet("user").(uuid.UUID)

	var user models.User
	if err := lib.DB.Select("id", "username").Where("id = ?", userId.String()).First(&user).Error; err != nil {
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
