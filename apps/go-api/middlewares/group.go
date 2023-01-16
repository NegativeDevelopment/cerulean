package middlewares

import (
	"github.com/NegativeDevelopment/cerulean/go-api/lib"
	"github.com/NegativeDevelopment/cerulean/go-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GroupOwnerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		groupId := c.Param("groupid")
		userId := c.MustGet("user").(uuid.UUID)

		var group models.Group
		if err := lib.DB.Where("id = ?", groupId).First(&group).Error; err != nil {
			c.JSON(400, gin.H{"error": "Group not found"})
			c.Abort()
			return
		}

		if group.OwnerID != userId {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func GroupMemberMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		groupId := c.Param("groupid")
		userId := c.MustGet("user").(uuid.UUID)

		var group models.Group
		if err := lib.DB.Preload("Members").Where("id = ?", groupId).First(&group).Error; err != nil {
			c.JSON(400, gin.H{"error": "Group not found"})
			c.Abort()
			return
		}

		if !group.IsMember(userId) {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
