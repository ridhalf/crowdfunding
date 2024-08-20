package service

import "crowdfunding/model/domain"

type CampaignService interface {
	FindAll() ([]domain.Campaign, error)
	FindByUserID(ID int) (domain.Campaign, error)
}
