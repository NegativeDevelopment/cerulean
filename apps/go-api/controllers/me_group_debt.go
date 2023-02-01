package controllers

import (
	"net/http"

	"github.com/NegativeDevelopment/cerulean/go-api/lib"
	"github.com/NegativeDevelopment/cerulean/go-api/middlewares"
	"github.com/NegativeDevelopment/cerulean/go-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func MyGroupDebtRoutes(parentGroup *gin.RouterGroup) {
	group := parentGroup.Group("/me/groups/:groupid/debts", middlewares.JwtAuthMiddleware(), middlewares.CheckUUIDMiddleware("groupid"), middlewares.GroupMemberMiddleware())
	{
		group.GET("", getGroupDebts)
		group.GET("/:userid", middlewares.CheckUUIDMiddleware("userid"), getDebtForUser)
	}
}

func getGroupDebts(c *gin.Context) {
	groupId := c.MustGet("groupid").(uuid.UUID)

	var debts []models.Debt
	if err := lib.DB.Preload(clause.Associations).Where("group_id = ?", groupId).Find(&debts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get debts"})
		return
	}

	if len(debts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Debts for group not found"})
		return
	}

	c.JSON(http.StatusOK, debts)
}

func getDebtForUser(c *gin.Context) {
	groupId := c.MustGet("groupid").(uuid.UUID)
	userId := c.MustGet("userid").(uuid.UUID)

	var debt []models.Debt
	if err := lib.DB.Preload(clause.Associations).Where("group_id = ? AND creditor_id = ?", groupId, userId).Find(&debt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get debts"})
		return
	}

	if len(debt) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Debts for user not found"})
		return
	}

	c.JSON(http.StatusOK, debt)
}
