package dto

type CouponGenerateRequest struct {
	CampaignID int32  `json:"campaign_id"`
	Quantity   int    `json:"quantity"`
	Prefix     string `json:"prefix"`
	Suffix     string `json:"suffix"`
	Length     int    `json:"length"`
}
