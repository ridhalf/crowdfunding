package service

import "crowdfunding/model/domain"

type CampaignService interface {
	FindByUserID(userID int) ([]domain.Campaign, error)
}
