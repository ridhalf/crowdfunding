package controller

import "github.com/gin-gonic/gin"

type CampaignController interface {
	FindAll(ctx *gin.Context)
}
