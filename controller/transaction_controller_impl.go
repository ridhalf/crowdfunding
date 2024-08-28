package controller

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
	"crowdfunding/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionControllerImpl struct {
	transactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{
		transactionService: transactionService,
	}
}

func (controller TransactionControllerImpl) FindByCampaignID(ctx *gin.Context) {
	var request web.TrasactionRequestByCampaignID
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		controller.failedTransactions(ctx, false)
		return
	}

	user := ctx.MustGet("user").(domain.User)
	request.User = user

	transactions, err := controller.transactionService.FindByCampaignID(request)
	if err != nil {
		controller.failedTransactions(ctx, true)
		return
	}
	response := web.ToTransactionResponseCampaigns(transactions)
	result := helper.Ok("list all transactions", response)
	ctx.JSON(http.StatusOK, result)
}
func (controller TransactionControllerImpl) FindByUserID(ctx *gin.Context) {
	user := ctx.MustGet("user").(domain.User)
	transactions, err := controller.transactionService.FindByUserID(user.ID)
	if err != nil {
		controller.failedTransactions(ctx, false)
		return
	}
	response := web.ToTransactionResponseUsers(transactions)
	result := helper.Ok("list all transactions", response)
	ctx.JSON(http.StatusOK, result)
}
func (controller TransactionControllerImpl) failedTransactions(ctx *gin.Context, forbidden bool) {
	if forbidden {
		response := helper.Forbidden("user is not the owner of the campaign", nil)
		ctx.JSON(http.StatusForbidden, response)
	} else {
		response := helper.BadRequest("failed to get campaign transaction", nil)
		ctx.JSON(http.StatusBadRequest, response)
	}

}
