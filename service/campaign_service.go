package service

import "crowdfunding/model/domain"

type CampaignService interface {
	FindCampaigns(userID int) ([]domain.Campaign, error)
}
