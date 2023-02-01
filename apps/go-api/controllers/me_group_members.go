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

	var members []models.Member
	if err := lib.DB.Where("group_id = ?", groupId).Find(&members).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get group members"})
		return
	}

	// Start transaction to add group member and create initial debts
	tx := lib.DB.Begin()
	// Create group member
	if err := tx.Create(&groupMember).Error; err != nil {
		if strings := strings.Contains(err.Error(), "duplicate key"); strings {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Group member already exists"})

			tx.Rollback()
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add group member"})
			tx.Rollback()
		}
	}

	// Create initial debts for every member
	for i := range members {
		dept := models.Debt{
			CreditorID: addGroupMemberInput.UserId,
			DebtorID:   members[i].UserID,
			GroupID:    groupId,
		}
		if err := tx.Create(&dept).Error; err != nil {
			if strings := strings.Contains(err.Error(), "duplicate key"); strings {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Initial debt with new member as creditor already exists"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create inital debt with new member as creditor"})
			}

			tx.Rollback()
			return
		}

		if err := lib.DB.Create(&models.Debt{
			CreditorID: members[i].UserID,
			DebtorID:   addGroupMemberInput.UserId,
			GroupID:    groupId,
		}).Error; err != nil {
			if strings := strings.Contains(err.Error(), "duplicate key"); strings {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Initial debt with new member as debtor already exists"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create inital debt with new member as debtor"})
			}

			tx.Rollback()
			return
		}
	}

	// Create initial debt with member as debtor and creditor
	if err := tx.Create(&models.Debt{
		CreditorID: addGroupMemberInput.UserId,
		DebtorID:   addGroupMemberInput.UserId,
		GroupID:    groupId,
	}).Error; err != nil {
		if strings := strings.Contains(err.Error(), "duplicate key"); strings {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Initial debt with new member as debtor an creditor already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create inital debt with new member as debtor and creditor"})
		}

		tx.Rollback()
		return
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add group member"})
		return
	}

	c.JSON(200, groupMember)
}
