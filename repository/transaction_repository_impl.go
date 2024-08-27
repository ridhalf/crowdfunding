package repository

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{
		db: db,
	}
}

func (repository TransactionRepositoryImpl) FindByCampaignID(campaignID int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := repository.db.Preload("User").Where("campaign_id = ?", campaignID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return helper.ResultOrError(transactions, err)
	}
	return transactions, nil
}
