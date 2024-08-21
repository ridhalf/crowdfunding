package controller

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
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

func (controller CampaignControllerImpl) FindAll(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Query("user_id"))
	campaigns, err := controller.campaignService.FindAll(userID)
	if err != nil {
		controller.failedGetCampaigns(ctx)
		return
	}
	result := helper.Ok("list all campaigns", web.ToCampaignsResponse(campaigns))
	ctx.JSON(http.StatusOK, result)
}

func (controller CampaignControllerImpl) FindByID(ctx *gin.Context) {
	var request web.CampaignRequestByID
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		controller.failedGetCampaigns(ctx)
		return
	}
	campaign, err := controller.campaignService.FindByID(request)
	if err != nil {
		controller.failedGetCampaigns(ctx)
		return
	}
	response := web.ToCampaignDetailResponse(campaign)
	result := helper.Ok("campaign detail", response)
	ctx.JSON(http.StatusOK, result)
}

func (controller CampaignControllerImpl) Create(ctx *gin.Context) {
	requestCampaign := web.CampaignRequestCreate{}
	err := ctx.ShouldBindJSON(&requestCampaign)
	if err != nil {
		result := helper.UnprocessableEntity("failed to create campaign", err)
		ctx.JSON(http.StatusUnprocessableEntity, result)
		return
	}
	user := ctx.MustGet("user").(domain.User)
	requestCampaign.User = user

	create, err := controller.campaignService.Create(requestCampaign)
	if err != nil {
		result := helper.UnprocessableEntity("failed to create campaign", err)
		ctx.JSON(http.StatusUnprocessableEntity, result)
	}
	response := web.ToCampaignResponse(create)
	result := helper.Ok("save campaign", response)
	ctx.JSON(http.StatusOK, result)
}

func (controller CampaignControllerImpl) failedGetCampaigns(ctx *gin.Context) {
	response := helper.BadRequest("error to get campaigns", nil)
	ctx.JSON(http.StatusBadRequest, response)
}
