package controllers

import (
	"net/http"
	"strings"

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
		group.POST("", middlewares.JwtAuthMiddleware(), createGroup)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group := models.Group{
		Name:    createGroupInput.Name,
		OwnerID: userId,
		Members: []models.Member{{UserID: userId}},
		Debts:   []models.Debt{{CreditorID: userId, DebtorID: userId}},
	}

	if err := lib.DB.Create(&group).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Group already exists"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create group"})
			return
		}
	}

	c.JSON(http.StatusCreated, group)
}
