package main

import (
	"crowdfunding/app"
	"crowdfunding/auth"
	"crowdfunding/controller"
	"crowdfunding/helper"
	"crowdfunding/repository"
	"crowdfunding/service"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	app.Env()
	db := app.NewDB()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserServiceImpl(userRepository)
	authJwt := auth.NewJwtService()
	userController := controller.NewUserController(userService, authJwt)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userController.Register)
	api.POST("/users/login", userController.Login)
	api.POST("/users/email_checker", userController.IsEmailAvailable)
	api.POST("/users/avatar", userController.UploadAvatar)

	err := router.Run("localhost:3000")
	helper.PanicIfError(err)
}
