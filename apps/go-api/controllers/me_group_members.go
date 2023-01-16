package controllers

import (
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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	groupMember := models.Member{
		GroupID: groupId,
		UserID:  addGroupMemberInput.UserId,
	}
	if err := lib.DB.Create(&groupMember).Error; err != nil {
		c.JSON(400, gin.H{"error": "Internal server error while creating group member"})
		return
	}

	c.JSON(200, groupMember)
}
