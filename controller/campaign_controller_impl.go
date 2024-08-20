package controller

import (
	"crowdfunding/helper"
	"crowdfunding/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CampaignControllerImpl struct {
	campaignService service.CampaignService
}

func NewCampaignController(campaignService service.CampaignService) CampaignController {
	return &CampaignControllerImpl{
		campaignService: campaignService,
	}
}

func (controller CampaignControllerImpl) FindCampaigns(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		controller.failedGetCampaigns(ctx)
		return
	}
	campaigns, err := controller.campaignService.FindCampaigns(userID)
	if err != nil {
		controller.failedGetCampaigns(ctx)
		return
	}
	response := helper.Ok("list all campaigns", campaigns)
	ctx.JSON(http.StatusOK, response)
}
func (controller CampaignControllerImpl) failedGetCampaigns(ctx *gin.Context) {
	response := helper.BadRequest("error to get campaigns", nil)
	ctx.JSON(http.StatusBadRequest, response)
}
