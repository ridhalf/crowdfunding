package service

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
	"crowdfunding/repository"
)

type CampaignServiceImpl struct {
	campaignRepository repository.CampaignRepository
}

func NewCampaignService(campaignRepository repository.CampaignRepository) CampaignService {
	return &CampaignServiceImpl{
		campaignRepository: campaignRepository,
	}
}

func (service CampaignServiceImpl) FindAll(userID int) ([]domain.Campaign, error) {
	if userID != 0 {
		campaigns, err := service.campaignRepository.FindByUserID(userID)
		return helper.ResultOrError(campaigns, err)
	}
	campaigns, err := service.campaignRepository.FindAll()
	return helper.ResultOrError(campaigns, err)
}

func (service CampaignServiceImpl) FindByID(request web.CampaignRequestByID) (*domain.Campaign, error) {
	campaign, err := service.campaignRepository.FindByID(request.ID)
	return helper.ResultOrError(campaign, err)
}
