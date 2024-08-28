package domain

import "time"

type Transaction struct {
	ID         int       `json:"id"`
	CampaignID int       `json:"campaign_id"`
	UserID     int       `json:"user_id"`
	Amount     int       `json:"amount"`
	Status     string    `json:"status"`
	Code       int       `json:"code"`
	User       User      `json:"user"`
	Campaign   Campaign  `json:"campaign"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
