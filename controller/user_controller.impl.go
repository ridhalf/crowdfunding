package controller

import (
	"crowdfunding/helper"
	"crowdfunding/model/web"
	"crowdfunding/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserControllerImpl(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (controller *UserControllerImpl) Register(ctx *gin.Context) {
	userCreateRequest := web.UserRequestRegister{}
	err := ctx.ShouldBindJSON(&userCreateRequest)
	if err != nil {
		response := helper.UnprocessableEntity("register account failed", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := controller.userService.RegisterUser(userCreateRequest)
	if err != nil {
		response := helper.BadRequest("register account failed")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := web.ToUserResponse(user, "token")
	result := helper.Ok("Account has been registered", response)
	ctx.JSON(http.StatusOK, result)
}
