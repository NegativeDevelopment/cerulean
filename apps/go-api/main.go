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
	lib.DB.AutoMigrate(&models.Group{})
	lib.DB.AutoMigrate(&models.Member{})
	lib.DB.AutoMigrate(&models.Transaction{})
	lib.DB.AutoMigrate(&models.TransactionMember{})
}

func main() {
	router := gin.Default()

	rootGroup := router.Group("/api")
	controllers.UsersRoutes(rootGroup)
	controllers.AuthRoutes(rootGroup)
	controllers.GroupsRoutes(rootGroup)

	controllers.MyRoutes(rootGroup)
	controllers.MyGroupsRoutes(rootGroup)
	controllers.MyGroupMembersRoutes(rootGroup)
	controllers.MyGroupTransactionsRoutes(rootGroup)

	router.Run()
}
