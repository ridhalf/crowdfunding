package service

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
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

func (service CampaignServiceImpl) FindCampaigns(userID int) ([]domain.Campaign, error) {
	if userID != 0 {
		campaigns, err := service.campaignRepository.FindByUserID(userID)
		return helper.ResultOrError(campaigns, err)
	}
	campaigns, err := service.campaignRepository.FindAll()
	return helper.ResultOrError(campaigns, err)
}
