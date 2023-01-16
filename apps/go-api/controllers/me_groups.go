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

func MyGroupsRoutes(parentGroup *gin.RouterGroup) {
	group := parentGroup.Group("/me/groups", middlewares.JwtAuthMiddleware())
	{
		group.GET("", getMyGroups)
		group.GET("/:groupid", middlewares.GroupMemberMiddleware(), getMyGroup)
	}
}

func getMyGroup(c *gin.Context) {
	groupId := c.Param("groupid")

	var group models.Group
	if err := lib.DB.Preload("Members").Where("id = ?", groupId).First(&group).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get group"})
			return
		}
	}

	c.JSON(http.StatusOK, group)
}

func getMyGroups(c *gin.Context) {
	userId := c.MustGet("user").(uuid.UUID)

	var groups []models.Group
	if err := lib.DB.Preload("Members", "user_id = ?", userId).Find(&groups).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get groups"})
		return
	}

	c.JSON(http.StatusOK, groups)
}
