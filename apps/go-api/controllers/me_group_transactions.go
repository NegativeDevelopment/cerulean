package controllers

import (
	"net/http"

	"github.com/NegativeDevelopment/cerulean/go-api/lib"
	"github.com/NegativeDevelopment/cerulean/go-api/middlewares"
	"github.com/NegativeDevelopment/cerulean/go-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Routes
func MyGroupTransactionsRoutes(r *gin.RouterGroup) {
	group := r.Group("/me/groups/:groupid/transactions",
		middlewares.JwtAuthMiddleware(), middlewares.CheckUUIDMiddleware("groupid"), middlewares.GroupMemberMiddleware())
	{
		group.GET("", getTransactions)
		group.POST("", createTransaction)
	}
}

// Handlers
func getTransactions(c *gin.Context) {
	groupId := c.MustGet("groupid").(uuid.UUID)

	var transactions []models.Transaction
	if err := lib.DB.Preload(clause.Associations).Where("group_id = ?", groupId).Find(&transactions).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transactions"})
			return
		}
	}

	c.JSON(http.StatusOK, transactions)
}

type createTransactionInput struct {
	Amount      float64  `json:"amount" binding:"required"`
	Currency    string   `json:"currency" binding:"required"`
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Members     []string `json:"members" binding:"required"`
}

func createTransaction(c *gin.Context) {
	userId := c.MustGet("user").(uuid.UUID)
	groupId := c.MustGet("groupid").(uuid.UUID)

	var input createTransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var members []models.TransactionMember
	for _, memberId := range input.Members {
		memberId, err := uuid.Parse(memberId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid member id"})
			return
		}

		members = append(members, models.TransactionMember{
			UserID: memberId,
		})
	}

	transaction := models.Transaction{
		GroupID:         groupId,
		Amount:          input.Amount,
		Currency:        input.Currency,
		Title:           input.Title,
		Description:     input.Description,
		CreatedByUserID: userId,
		Members:         members,
	}

	if err := lib.DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	if err := lib.DB.Preload(clause.Associations).First(&transaction, transaction.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transaction"})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}
