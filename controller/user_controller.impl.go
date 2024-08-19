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

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (controller *UserControllerImpl) Register(ctx *gin.Context) {
	registerRequest := web.UserRequestRegister{}
	err := ctx.ShouldBindJSON(&registerRequest)
	if err != nil {
		response := helper.UnprocessableEntity("register account failed", err)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := controller.userService.Register(registerRequest)
	if err != nil {
		response := helper.BadRequest("register account failed")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := web.ToUserResponse(user, "token")
	result := helper.Ok("Account has been registered", response)
	ctx.JSON(http.StatusOK, result)
}

func (controller *UserControllerImpl) Login(ctx *gin.Context) {
	//TODO implement me
	loginRequest := web.UserRequestLogin{}
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		response := helper.UnprocessableEntity("login account failed", err)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	login, err := controller.userService.Login(loginRequest)
	if err != nil {
		response := helper.UnprocessableEntityString("login account failed", err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := web.ToUserResponse(login, "token")
	result := helper.Ok(`login successful. welcome back!`, response)
	ctx.JSON(http.StatusOK, result)
}

func (controller *UserControllerImpl) IsEmailAvailable(ctx *gin.Context) {

	emailCheck := web.UserRequestEmailCheck{}
	err := ctx.ShouldBindJSON(&emailCheck)
	if err != nil {
		response := helper.UnprocessableEntity("email check failed", err)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	available, err := controller.userService.IsEmailAvailable(emailCheck)
	if err != nil {
		response := helper.UnprocessableEntityString("an unexpected error occurred. Please try again later", err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data := gin.H{"is_available": available}
	var metaString string
	if available {
		metaString = "email is available"
	} else {
		metaString = "email is not available"
	}
	result := helper.Ok(metaString, data)
	ctx.JSON(http.StatusOK, result)
}
