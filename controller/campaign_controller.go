package controller

import "github.com/gin-gonic/gin"

type CampaignController interface {
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Create(ctx *gin.Context)
}
