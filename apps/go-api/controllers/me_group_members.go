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
func MyGroupMembersRoutes(parentGroup *gin.RouterGroup) {
	group := parentGroup.Group("me/groups/:groupid/members", middlewares.JwtAuthMiddleware(), middlewares.CheckUUIDMiddleware("groupid"))
	{
		group.POST("", middlewares.GroupOwnerMiddleware(), addGroupMember)
	}
}

// --- Handlers ---
type AddGroupMemberInput struct {
	UserId uuid.UUID `json:"userid" binding:"required"`
}

func addGroupMember(c *gin.Context) {
	var groupId = c.MustGet("groupid").(uuid.UUID)

	var addGroupMemberInput AddGroupMemberInput
	if err := c.ShouldBindJSON(&addGroupMemberInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupMember := models.Member{
		GroupID: groupId,
		UserID:  addGroupMemberInput.UserId,
	}
	if err := lib.DB.Create(&groupMember).Error; err != nil {
		if strings := strings.Contains(err.Error(), "duplicate key"); strings {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Group member already exists"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add group member"})
			return
		}
	}

	c.JSON(200, groupMember)
}
