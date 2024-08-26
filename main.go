package main

import (
	"crowdfunding/app"
	"crowdfunding/auth"
	"crowdfunding/controller"
	"crowdfunding/helper"
	"crowdfunding/middleware"
	"crowdfunding/repository"
	"crowdfunding/service"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	app.Env()
	db := app.NewDB()

	//repositories
	userRepository := repository.NewUserRepository(db)
	campaignRepository := repository.NewCampaignRepository(db)
	//services
	userService := service.NewUserServiceImpl(userRepository)
	campaignService := service.NewCampaignService(campaignRepository)
	//middleware
	authJwt := auth.NewJwtService()
	authMiddleware := middleware.AuthMiddleware(authJwt, userService)

	//controllers
	userController := controller.NewUserController(userService, authJwt)
	campaignController := controller.NewCampaignController(campaignService)

	router := gin.Default()
	router.Static("/images", "./images")

	api := router.Group("/api/v1")
	api.POST("/users", userController.Register)
	api.POST("/users/login", userController.Login)
	api.POST("/users/email_checker", userController.IsEmailAvailable)
	api.POST("/users/avatar", authMiddleware, userController.UploadAvatar)

	api.GET("/campaigns", campaignController.FindAll)
	api.GET("/campaigns/:id", campaignController.FindByID)
	api.POST("/campaigns", authMiddleware, campaignController.Create)
	api.PUT("/campaigns/:id", authMiddleware, campaignController.Update)

	err := router.Run("localhost:3000")
	helper.PanicIfError(err)
}
