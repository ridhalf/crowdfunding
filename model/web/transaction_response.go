package web

import "crowdfunding/model/domain"

type TransactionResponseCampaign struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Amount    int    `json:"amount"`
	CreatedAt string `json:"created_at"`
}

func ToTransactionResponseCampaign(transaction domain.Transaction) TransactionResponseCampaign {
	return TransactionResponseCampaign{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt.String(),
	}
}
func ToTransactionResponseCampaigns(transactions []domain.Transaction) []TransactionResponseCampaign {
	if len(transactions) == 0 {
		return []TransactionResponseCampaign{}
	}
	var transactionResponses []TransactionResponseCampaign
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, ToTransactionResponseCampaign(transaction))
	}
	return transactionResponses
}
