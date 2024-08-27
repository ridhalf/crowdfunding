package controller

import "github.com/gin-gonic/gin"

type TransactionController interface {
	FindByCampaignID(ctx *gin.Context)
}
