package test

import (
	"crowdfunding/app"
	"crowdfunding/helper"
	"crowdfunding/model/web"
	"crowdfunding/repository"
	"crowdfunding/service"
	"testing"
)

func TestUserCreate(t *testing.T) {
	db := app.NewDB()
	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository)
	user := web.UserCreateRequest{
		Name:       "John Doe",
		Occupation: "Programmer",
		Email:      "john.doe@gmail.com",
		Password:   "password",
	}
	_, err := userService.RegisterUser(user)
	helper.PanicIfError(err)
}
