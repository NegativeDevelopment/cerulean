package main

import (
	"github.com/NegativeDevelopment/cerulean/go-api/controllers"
	"github.com/NegativeDevelopment/cerulean/go-api/lib"
	"github.com/NegativeDevelopment/cerulean/go-api/models"
	"github.com/gin-gonic/gin"
)

func init() {
	lib.InitializeDotEnv()
	lib.InitializeDatabase()

	lib.DB.AutoMigrate(&models.User{})
}

func main() {
	router := gin.Default()

	rootGroup := router.Group("/api")
	controllers.UserRoutes(rootGroup)
	controllers.AuthRoutes(rootGroup)
	controllers.MeRoutes(rootGroup)

	router.Run()
}
