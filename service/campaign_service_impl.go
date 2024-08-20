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

func (service CampaignServiceImpl) FindAll() ([]domain.Campaign, error) {
	campaigns, err := service.campaignRepository.FindAll()
	return helper.ResultOrError(campaigns, err)
}

func (service CampaignServiceImpl) FindById(ID int) (domain.Campaign, error) {
	//TODO implement me
	panic("implement me")
}
