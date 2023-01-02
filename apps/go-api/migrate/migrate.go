package main

import (
	"github.com/NegativeDevelopment/cerulean/go-api/initializers"
	"github.com/NegativeDevelopment/cerulean/go-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
