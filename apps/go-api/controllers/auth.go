package controllers

import (
	"github.com/NegativeDevelopment/cerulean/go-api/lib"
	"github.com/NegativeDevelopment/cerulean/go-api/models"
	"github.com/gin-gonic/gin"
)

// --- Routes ---
func AuthRoutes(parentGroup *gin.RouterGroup) {
	group := parentGroup.Group("/auth")
	{
		group.POST("/login", login)
		group.POST("/register", register)
	}
}

// --- Handlers ---
type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func login(c *gin.Context) {
	var login Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := lib.DB.Where("username = ?", login.Username).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "Invalid username or password"})
		return
	}

	if !user.CheckPassword(login.Password) {
		c.JSON(400, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := lib.GenerateToken(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate jwt token"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func register(c *gin.Context) {
	var register Register
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	if err := lib.DB.Where("username = ?", register.Username).First(&existingUser).Error; err == nil {
		c.JSON(400, gin.H{"error": "Username already taken"})
		return
	}

	var newUser models.User
	newUser.Username = register.Username
	newUser.Password = register.Password
	newUser.HashPassword()

	if err := lib.DB.Create(&newUser).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(200, gin.H{"message": "User created"})
}
