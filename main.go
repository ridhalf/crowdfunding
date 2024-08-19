package main

import (
	"crowdfunding/app"
	"crowdfunding/controller"
	"crowdfunding/helper"
	"crowdfunding/repository"
	"crowdfunding/service"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	db := app.NewDB()
	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository)
	userController := controller.NewUserControllerImpl(userService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userController.Register)

	err := router.Run("localhost:3000")
	helper.PanicIfError(err)
}
