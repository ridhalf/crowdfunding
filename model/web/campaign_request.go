package web

import "crowdfunding/model/domain"

type CampaignRequestByID struct {
	ID int `uri:"id" binding:"required"`
}
type CampaignRequestCreate struct {
	Name             string      `json:"name" binding:"required"`
	ShortDescription string      `json:"short_description" binding:"required"`
	Description      string      `json:"description" binding:"required"`
	GoalAmount       int         `json:"goal_amount" binding:"required"`
	Perks            string      `json:"perks" binding:"required"`
	User             domain.User `json:"user"`
}
