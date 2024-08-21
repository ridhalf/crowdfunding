package service

import (
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
)

type CampaignService interface {
	FindAll(userID int) ([]domain.Campaign, error)
	FindByID(request web.CampaignRequestByID) (domain.Campaign, error)
	Create(request web.CampaignRequestCreate) (domain.Campaign, error)
}
