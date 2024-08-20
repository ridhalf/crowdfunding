package repository

import "crowdfunding/model/domain"

type CampaignRepository interface {
	FindAll() ([]domain.Campaign, error)
	FindByUserID(userID int) (domain.Campaign, error)
}
