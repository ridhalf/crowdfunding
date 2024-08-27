package service

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
	"crowdfunding/repository"
	"errors"
)

type TransactionServiceImpl struct {
	transactionRepository repository.TransactionRepository
	campaignRepository    repository.CampaignRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository, campaignRepository repository.CampaignRepository) TransactionService {
	return &TransactionServiceImpl{
		transactionRepository: transactionRepository,
		campaignRepository:    campaignRepository,
	}
}

func (service TransactionServiceImpl) FindByCampaignID(request web.TrasactionRequestByCampaignID) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	campaign, err := service.campaignRepository.FindByID(request.CampaignID)
	if err != nil {
		return helper.ResultOrError(transactions, err)
	}

	if campaign.UserID != request.User.ID {
		return helper.ResultOrError(transactions, errors.New("user id is not match"))
	}

	transactions, err = service.transactionRepository.FindByCampaignID(request.CampaignID)
	if err != nil {
		return helper.ResultOrError(transactions, err)
	}
	return transactions, nil
}
