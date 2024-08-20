package web

type CampaignRequestByID struct {
	ID int `uri:"id" binding:"required"`
}
