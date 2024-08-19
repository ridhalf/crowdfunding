package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	RegisterUser(ctx *gin.Context)
}
