package web

import "crowdfunding/model/domain"

type TrasactionRequestByCampaignID struct {
	CampaignID int         `uri:"id" binding:"required"`
	User       domain.User `json:"user"`
}
