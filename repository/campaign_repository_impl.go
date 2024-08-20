package repository

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"gorm.io/gorm"
)

type CampaignRepositoryImpl struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) CampaignRepository {
	return &CampaignRepositoryImpl{
		db: db,
	}
}

func (repository CampaignRepositoryImpl) FindAll() ([]domain.Campaign, error) {
	var campaigns []domain.Campaign
	err := repository.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	return helper.ResultOrError(campaigns, err)
}

func (repository CampaignRepositoryImpl) FindByUserID(userID int) ([]domain.Campaign, error) {
	var campaigns []domain.Campaign
	err := repository.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns, "user_id = ?", userID).Error
	return helper.ResultOrError(campaigns, err)

}

func (repository CampaignRepositoryImpl) FindByID(ID int) (*domain.Campaign, error) {
	var campaign *domain.Campaign
	err := repository.db.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error
	return helper.ResultOrError(campaign, err)
}
