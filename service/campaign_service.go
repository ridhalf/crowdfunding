package service

import "crowdfunding/model/domain"

type CampaignService interface {
	FindAll(userID int) ([]domain.Campaign, error)
}
