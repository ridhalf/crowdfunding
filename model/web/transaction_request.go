package web

import (
	"crowdfunding/model/domain"
)

type TrasactionRequestByCampaignID struct {
	CampaignID int         `uri:"id" binding:"required"`
	User       domain.User `json:"user"`
}
type TransactionRequestCreate struct {
	Amount     int         `json:"amount" binding:"required"`
	CampaignID int         `json:"campaign_id" binding:"required"`
	User       domain.User `json:"user"`
}
