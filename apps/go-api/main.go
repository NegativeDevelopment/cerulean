package main

import (
	"fmt"

	"github.com/NegativeDevelopment/cerulean/go-api/controllers"
	"github.com/NegativeDevelopment/cerulean/go-api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	fmt.Println("Hello World")

	router := gin.Default()

	rootGroup := router.Group("/api")
	{
		rootGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	controllers.UserRoutes(rootGroup)
	controllers.AuthRoutes(rootGroup)

	router.Run()
}
