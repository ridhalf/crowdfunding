package repository

import "crowdfunding/model/domain"

type TransactionRepository interface {
	FindByCampaignID(campaignID int) ([]domain.Transaction, error)
}
