package service

import (
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
)

type TransactionService interface {
	FindByCampaignID(request web.TrasactionRequestByCampaignID) ([]domain.Transaction, error)
	FindByUserID(userID int) ([]domain.Transaction, error)
	Create(request web.TransactionRequestCreate) (domain.Transaction, error)
}
