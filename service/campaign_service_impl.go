package service

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
	"crowdfunding/repository"
	"github.com/gosimple/slug"
	"strconv"
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

func (service CampaignServiceImpl) FindByID(request web.CampaignRequestByID) (domain.Campaign, error) {
	campaign, err := service.campaignRepository.FindByID(request.ID)
	return helper.ResultOrError(campaign, err)
}

func (service CampaignServiceImpl) Create(request web.CampaignRequestCreate) (domain.Campaign, error) {
	slug.Make(request.Name + strconv.Itoa(request.User.ID))
	campaign := domain.Campaign{
		Name:             request.Name,
		ShortDescription: request.ShortDescription,
		Description:      request.Description,
		GoalAmount:       request.GoalAmount,
		Perks:            request.Perks,
		UserID:           request.User.ID,
		Slug:             slug.Make(request.Name + " " + strconv.Itoa(request.User.ID)),
	}
	result, err := service.campaignRepository.Save(campaign)
	return helper.ResultOrError(result, err)

}
