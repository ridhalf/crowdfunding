package service

import "crowdfunding/model/domain"

type CampaignService interface {
	FindAll() ([]domain.Campaign, error)
	FindById(ID int) (domain.Campaign, error)
}
