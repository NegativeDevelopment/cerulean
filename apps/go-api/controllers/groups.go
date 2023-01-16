package controllers

import (
	"github.com/NegativeDevelopment/cerulean/go-api/lib"
	"github.com/NegativeDevelopment/cerulean/go-api/middlewares"
	"github.com/NegativeDevelopment/cerulean/go-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// --- Routes ---
func GroupsRoutes(parentGroup *gin.RouterGroup) {
	group := parentGroup.Group("/groups")
	{
		group.POST("/", middlewares.JwtAuthMiddleware(), createGroup)
	}
}

// --- Handlers ---
type CreateGroupInput struct {
	Name string `json:"name" binding:"required"`
}

func createGroup(c *gin.Context) {
	userId := c.MustGet("user").(uuid.UUID)

	var createGroupInput CreateGroupInput
	if err := c.ShouldBindJSON(&createGroupInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	group := models.Group{
		Name:    createGroupInput.Name,
		OwnerID: userId,
		Members: []models.Member{{UserID: userId}},
	}

	if err := lib.DB.Create(&group).Error; err != nil {
		c.JSON(400, gin.H{"error": "Internal server error while creating group"})
		return
	}

	c.JSON(200, group)
}
